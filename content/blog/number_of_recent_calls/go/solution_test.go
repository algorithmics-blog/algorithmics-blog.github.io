package number_of_recent_calls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RecentCounter(t *testing.T) {
	t.Run("[1], [100], [3001], [3002]", func(t *testing.T) {
		counter := Constructor()

		counter.Ping(1)
		counter.Ping(100)
		counter.Ping(3001)
		counter.Ping(3002)
		res := counter.Ping(3003)

		assert.Equal(t, 4, res)
	})

	t.Run("[1], [100], [3001]", func(t *testing.T) {
		counter := Constructor()

		counter.Ping(1)
		counter.Ping(100)
		res := counter.Ping(3001)

		assert.Equal(t, 3, res)
	})
}
