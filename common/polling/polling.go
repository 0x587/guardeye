// Package polling 轮询逻辑
package polling

import (
	"fmt"
	"sync"
)

type IF[Req, Rsp any] interface {
	Accept(key string, stream stream[Req, Rsp], killFunc func())
	Send(key string, value Req) error
}

func New[Req, Rsp any](callback func(key string, rsp Rsp)) IF[Req, Rsp] {
	return &impl[Req, Rsp]{
		conns:    make(map[string]*connPool[Req, Rsp]),
		callback: callback,
	}
}

type impl[Req, Rsp any] struct {
	sync.Mutex
	conns    map[string]*connPool[Req, Rsp]
	callback func(key string, rsp Rsp)
}

func (i *impl[Req, Rsp]) Accept(key string, s stream[Req, Rsp], killFunc func()) {
	p, ok := i.conns[key]
	if !ok {
		i.Lock()
		defer i.Unlock()
		i.conns[key] = &connPool[Req, Rsp]{
			streams: make(map[stream[Req, Rsp]]func()),
			callback: func(rsp Rsp) {
				i.callback(key, rsp)
			},
		}
		p = i.conns[key]
	}
	p.add(s, killFunc)
}

func (i *impl[Req, Rsp]) Send(key string, value Req) error {
	p, ok := i.conns[key]
	if !ok {
		return fmt.Errorf("no such key: %s", key)
	}
	return p.send(value)
}

type connPool[Req, Rsp any] struct {
	sync.Mutex
	//streams []lo.Tuple2[stream[T], func()]
	streams  map[stream[Req, Rsp]]func()
	callback func(Rsp)
}

func (p *connPool[Req, Rsp]) add(s stream[Req, Rsp], killFunc func()) {
	p.Lock()
	defer p.Unlock()
	p.streams[s] = killFunc
	go func() {
		for {
			rsp, err := s.Recv()
			if err != nil {
				killFunc()
				delete(p.streams, s)
				return
			}
			if err == nil {
				p.callback(rsp)
			}
		}
	}()
}

func (p *connPool[Req, Rsp]) send(v Req) error {
	p.Lock()
	defer p.Unlock()
	var err error
	for s, killFunc := range p.streams {
		err = s.Send(v)
		if err == nil {
			return nil
		}
		delete(p.streams, s)
		killFunc()
		return err
	}
	return nil
}

type stream[Req, Rsp any] interface {
	Send(Req) error
	Recv() (Rsp, error)
}
