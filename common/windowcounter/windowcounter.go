package windowcounter

import (
	"sync/atomic"
	"time"
)

type IF interface {
	Add(int)
	Get() int
}

func New(d time.Duration) IF {
	res := &impl{
		ticker: time.NewTicker(d),
	}
	go res.loop()
	return res
}

type impl struct {
	v      atomic.Uint64
	ticker *time.Ticker
}

func (i *impl) Add(v int) {
	i.v.Add(uint64(v))
}

func (i *impl) Get() int {
	return int(i.v.Load())
}

func (i *impl) loop() {
	for {
		select {
		case <-i.ticker.C:
			i.v.Store(0)
		}
	}
}
