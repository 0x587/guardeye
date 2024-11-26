// Package pool 对象池
package pool

import "sync/atomic"

type UseOption struct {
	NeedDestroy bool
}

type IF[T any] interface {
	Use(func(T) *UseOption)
}

func New[T any](size, lifespan int32, maker func() T, cleaner func(T)) IF[T] {
	ch := make(chan objInfo[T], size)
	go func() {
		for {
			ch <- objInfo[T]{
				obj:  maker(),
				life: &atomic.Int32{},
			}
		}
	}()
	return &impl[T]{
		objCh:    ch,
		cleaner:  cleaner,
		lifespan: lifespan,
	}
}

type impl[T any] struct {
	lifespan int32
	objCh    chan objInfo[T]
	cleaner  func(T)
}

type objInfo[T any] struct {
	obj  T
	life *atomic.Int32
}

func (i *impl[T]) Use(f func(T) *UseOption) {
	info := <-i.objCh
	o := f(info.obj)
	if o != nil {
		if o.NeedDestroy {
			i.cleaner(info.obj)
			return
		}
	}
	go func() {
		if i.lifespan > 0 && info.life.Load() == i.lifespan-1 {
			i.cleaner(info.obj)
		} else {
			info.life.Add(1)
			i.objCh <- info
		}
	}()
}
