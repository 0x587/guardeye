package limiter

import (
	"sync"
	"time"
)

type IF interface {
	Do(func() error)
}

func New(limit time.Duration) IF {
	return &impl{
		limit: limit,
	}
}

type impl struct {
	sync.Mutex
	lastRun time.Time
	limit   time.Duration
}

func (i *impl) Do(f func() error) {
	i.Lock()
	defer i.Unlock()
	if time.Since(i.lastRun) < i.limit {
		return
	}
	if err := f(); err == nil {
		i.lastRun = time.Now()
	}
}
