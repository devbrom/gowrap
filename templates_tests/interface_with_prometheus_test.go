package templatestests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestInterfaceWithPrometheus_F(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		impl := &testImpl{r1: "1", r2: "2"}

		wrapped := NewTestInterfaceWithPrometheus(impl, "test")
		r1, r2, err := wrapped.F(context.Background(), "a1", "a2")

		assert.NoError(t, err)
		assert.Equal(t, "1", r1)
		assert.Equal(t, "2", r2)
	})

	t.Run("with error", func(t *testing.T) {
		impl := &testImpl{r1: "1", r2: "2", err: errors.New("unexpected error")}

		wrapped := NewTestInterfaceWithPrometheus(impl, "test")
		r1, r2, err := wrapped.F(context.Background(), "a1", "a2")

		assert.Error(t, err)
		assert.Equal(t, "1", r1)
		assert.Equal(t, "2", r2)
	})
}
