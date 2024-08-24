package throttle

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := make([]int, 20)

	start := time.Now()
	for i, val := range Throttle(s, 3, time.Second) {
		fmt.Println(i, val)
		if i == 10 {
			break
		}
	}
	end := time.Now()
	assert.Equal(t, float64(3), end.Sub(start).Truncate(time.Second).Seconds())
}
