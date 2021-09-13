package optional

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func conventionalIntReturn(v int, err error) (int, error) {
	return v, err
}

func conventionalIntReturnPtr(v *int, err error) (*int, error) {
	return v, err
}

func TestResultInt(t *testing.T) {
	t.Run("NewResultInt", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			r := NewResultInt(conventionalIntReturn(1, nil))
			assert.Equal(t, 1, r.Value())
			assert.Panics(t, func() { _ = r.Err() })
		})
		t.Run("Err", func(t *testing.T) {
			r := NewResultInt(conventionalIntReturn(0, errors.New("")))
			assert.Panics(t, func() { _ = r.Value() })
			assert.Error(t, r.Err())
		})
	})
	t.Run("NewResultIntPtr", func(t *testing.T) {
		t.Run("Panic", func(t *testing.T) {
			assert.Panics(t, func() { _ = NewResultIntPtr(conventionalIntReturnPtr(nil, nil)) })
		})
		t.Run("Ok", func(t *testing.T) {
			v := 1
			r := NewResultIntPtr(conventionalIntReturnPtr(&v, nil))
			assert.Equal(t, 1, r.Value())
			assert.Panics(t, func() { _ = r.Err() })
		})
		t.Run("Err", func(t *testing.T) {
			r := NewResultIntPtr(conventionalIntReturnPtr(nil, errors.New("")))
			assert.Panics(t, func() { _ = r.Value() })
			assert.Error(t, r.Err())
		})
	})
	t.Run("Ok", func(t *testing.T) {
		r := OkInt(1)
		assert.Equal(t, 1, r.Value())
		assert.Panics(t, func() { _ = r.Err() })
		assert.True(t, r.Present())
		v, err := r.Get()
		assert.Equal(t, 1, v)
		assert.NoError(t, err)
	})
	t.Run("Err", func(t *testing.T) {
		r := ErrInt(errors.New(""))
		assert.Panics(t, func() { _ = r.Value() })
		assert.Error(t, r.Err())
		assert.False(t, r.Present())
		v, err := r.Get()
		assert.Equal(t, 0, v)
		assert.Error(t, err)
	})
}
