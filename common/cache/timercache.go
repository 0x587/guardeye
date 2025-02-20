package cache

import (
	"context"
	"sync"
	"time"

	"github.com/samber/lo"
)

func TimerCacheFunc[T any](f func(context.Context) (T, error), d time.Duration) func(context.Context) (T, error) {
	mutex := sync.Mutex{}
	updateAt := (*time.Time)(nil)
	var v T
	return func(ctx context.Context) (T, error) {
		mutex.Lock()
		defer mutex.Unlock()
		if updateAt == nil || time.Now().After((*updateAt).Add(d)) {
			newV, err := f(ctx)
			if err != nil {
				return *lo.Nil[T](), err
			}
			now := time.Now()
			updateAt, v = &now, newV
			return v, nil
		}
		return v, nil
	}
}
