package lifepool

import (
	"sync"
	"time"
)

func NewLifePool[T any, K comparable](lifespan time.Duration,
	maker func(K) T, cleaner func(T)) *LifePool[T, K] {
	return &LifePool[T, K]{
		lifespan: lifespan,
		maker:    maker,
		cleaner:  cleaner,
		objs:     make(map[K]objInfo[T]),
	}
}

type LifePool[T any, K comparable] struct {
	sync.RWMutex
	lifespan time.Duration
	maker    func(K) T
	cleaner  func(T)
	objs     map[K]objInfo[T]
}

type objInfo[T any] struct {
	obj    T
	deadAt time.Time
}

func (p *LifePool[T, K]) Get(key K) T {
	p.RLock()
	obj, ok := p.objs[key]
	p.RUnlock()
	if ok {
		if obj.deadAt.After(time.Now()) {
			obj.deadAt = time.Now().Add(p.lifespan)
			return obj.obj
		}
		p.cleaner(obj.obj)
	}
	o := p.maker(key)
	p.Lock()
	p.objs[key] = objInfo[T]{obj: o, deadAt: time.Now().Add(p.lifespan)}
	p.Unlock()
	return o
}
