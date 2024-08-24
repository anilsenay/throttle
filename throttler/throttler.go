package throttler

import (
	"sync"
	"time"
)

type Throttler struct {
	bucket      int
	bucketLimit int
	interval    time.Duration
	ticker      *time.Ticker
	done        bool
	waitCh      chan struct{}
	once        sync.Once
}

func New(ops int, interval time.Duration) *Throttler {
	return &Throttler{
		bucketLimit: ops,
		interval:    interval,
		waitCh:      make(chan struct{}, 1),
	}
}

func (t *Throttler) IsDone() bool {
	return t.done
}

func (t *Throttler) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
	}
	close(t.waitCh)
	t.done = true
}

func (t *Throttler) start() {
	t.ticker = time.NewTicker(t.interval)

	go func() {
		for range t.ticker.C {
			if t.bucket >= t.bucketLimit {
				t.bucket = 0
				t.waitCh <- struct{}{}
			}
		}
	}()
}

func (th *Throttler) Allow() bool {
	th.once.Do(th.start)

	if th.bucket >= th.bucketLimit {
		<-th.waitCh
	}

	if th.done {
		return false
	}

	th.bucket++
	return true
}
