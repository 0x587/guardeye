// Package polling 轮询逻辑
package polling

import (
	"errors"
	"fmt"
	"sync"
)

type IF[Req, Rsp any] interface {
	Accept(key string, stream stream[Req, Rsp])
	Send(key string, value Req) error
}

func New[Req, Rsp any](callback func(key string, rsp Rsp)) IF[Req, Rsp] {
	return &impl[Req, Rsp]{
		callback: callback,
		streams:  make(map[string]map[stream[Req, Rsp]]bool),
	}
}

type impl[Req, Rsp any] struct {
	sync.Mutex
	callback func(key string, rsp Rsp)
	streams  map[string]map[stream[Req, Rsp]]bool
}

func (i *impl[Req, Rsp]) Accept(key string, s stream[Req, Rsp]) {
	i.Lock()
	if i.streams[key] == nil {
		i.streams[key] = make(map[stream[Req, Rsp]]bool)
	}
	i.streams[key][s] = true
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
	}
}

func (i *impl[Req, Rsp]) Send(key string, value Req) error {
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

type stream[Req, Rsp any] interface {
	Send(Req) error
	Recv() (Rsp, error)
}
