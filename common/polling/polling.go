// Package polling 轮询逻辑
package polling

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type IF[Req, Rsp any] interface {
	Accept(key string, stream stream[Req, Rsp])
	Send(key string, value Req) error
	List() map[string]time.Time
}

type stream[Req, Rsp any] interface {
	Send(Req) error
	Recv() (Rsp, error)
}

func New[Req, Rsp any](callback func(key string, rsp Rsp)) IF[Req, Rsp] {
	return &impl[Req, Rsp]{
		callback: callback,
		streams:  make(map[string]map[stream[Req, Rsp]]bool),
		lastSeen: make(map[string]time.Time),
	}
}

type impl[Req, Rsp any] struct {
	sync.Mutex
	callback func(key string, rsp Rsp)
	streams  map[string]map[stream[Req, Rsp]]bool
	lastSeen map[string]time.Time
}

func (i *impl[Req, Rsp]) Accept(key string, s stream[Req, Rsp]) {
	i.Lock()
	if i.streams[key] == nil {
		i.streams[key] = make(map[stream[Req, Rsp]]bool)
	}
	i.streams[key][s] = true
	i.lastSeen[key] = time.Now()
	i.Unlock()
	for {
		recv, err := s.Recv()
		if err != nil {
			i.Lock()
			delete(i.streams[key], s)
			i.Unlock()
			return
		}
		i.callback(key, recv)
		i.Lock()
		i.lastSeen[key] = time.Now()
		i.Unlock()
	}
}

func (i *impl[Req, Rsp]) Send(key string, value Req) error {
	i.Lock()
	defer i.Unlock()
	streams := i.streams[key]
	if len(streams) == 0 {
		return errors.New(fmt.Sprintf("no suck link key: %s", key))
	}
	for s := range streams {
		err := s.Send(value)
		if err == nil {
			return nil
		}
	}
	return errors.New("no link could use")
}

func (i *impl[Req, Rsp]) List() map[string]time.Time {
	res := make(map[string]time.Time)
	i.Lock()
	defer i.Unlock()
	for k, v := range i.lastSeen {
		res[k] = v
	}
	return res
}
