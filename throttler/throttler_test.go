package throttler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	th := New(3, time.Second)
	assert.NotNil(t, th)
	assert.Equal(t, 3, th.bucketLimit)
	assert.Equal(t, time.Second, th.interval)

	start := time.Now()
	i := 0
	for th.Allow() {
		if i == 10 {
			th.Stop()
		}
		i++
	}
	end := time.Now()
	assert.Equal(t, float64(3), end.Sub(start).Truncate(time.Second).Seconds())

	assert.True(t, th.IsDone())
}
