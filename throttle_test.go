package throttle_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/anilsenay/throttle"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := make([]int, 20)

	start := time.Now()
	for i, val := range throttle.Limit(s, 3, time.Second) {
		fmt.Println(i, val)
		if i == 10 {
			break
		}
	}
	end := time.Now()
	assert.Equal(t, float64(3), end.Sub(start).Truncate(time.Second).Seconds())
}
