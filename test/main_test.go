package test

import (
	"github.com/elliotchance/testify-stats"
	"github.com/elliotchance/testify-stats/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(testify_stats.Run(m))
}

func TestFoo(t *testing.T) {
	assert.Equal(t, "foo", "foo")

	t.Run("Bar", func(t *testing.T) {
		assert.Equal(t, "foo", "foo")
	})
}

func TestBar(t *testing.T) {
	assert.Equal(t, "foo", "foo")
	assert.Equal(t, "foo", "foo")
}
