package throttle

import (
	"iter"
	"time"

	"github.com/anilsenay/throttle/throttler"
)

func Limit[Slice ~[]E, E any](s Slice, ops int, interval time.Duration) iter.Seq2[int, E] {
	th := throttler.New(ops, interval)

	return func(yield func(int, E) bool) {
		defer th.Stop()

		for i := 0; i < len(s); i++ {
			if th.Allow() && !yield(i, s[i]) {
				return
			}
		}
	}
}
