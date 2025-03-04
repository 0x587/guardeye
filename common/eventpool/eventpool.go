package eventpool

import (
	"sync"

	"golang.org/x/net/context"
)

type IF[K comparable, V any] interface {
	Invoke(k K, v V)
	Wait(ctx context.Context, k K) (V, error)
}

func New[K comparable, V any]() IF[K, V] {
	return &impl[K, V]{}
}

type impl[K comparable, V any] struct {
	sync.Mutex
	waiters sync.Map
}

func (i *impl[K, V]) Invoke(k K, v V) {
	i.Lock()
	defer i.Unlock()

	if chVal, ok := i.waiters.Load(k); ok {
		if chans, ok := chVal.([]chan V); ok {
			for _, ch := range chans {
				ch <- v
				close(ch)
			}
			i.waiters.Delete(k)
		}
	}
}

func (i *impl[K, V]) Wait(ctx context.Context, k K) (V, error) {
	ch := make(chan V, 1)

	i.Lock()
	if chVal, ok := i.waiters.Load(k); ok {
		if chans, ok := chVal.([]chan V); ok {
			i.waiters.Store(k, append(chans, ch))
		}
	} else {
		i.waiters.Store(k, []chan V{ch})
	}
	i.Unlock()

	select {
	case val := <-ch:
		return val, nil
	case <-ctx.Done():
		i.Lock()
		if chVal, ok := i.waiters.Load(k); ok {
			if chans, ok := chVal.([]chan V); ok {
				// 移除超时的 channel
				newChans := make([]chan V, 0, len(chans))
				for _, c := range chans {
					if c != ch {
						newChans = append(newChans, c)
					}
				}
				if len(newChans) == 0 {
					i.waiters.Delete(k)
				} else {
					i.waiters.Store(k, newChans)
				}
			}
		}
		i.Unlock()
		var zero V
		return zero, ctx.Err()
	}
}
