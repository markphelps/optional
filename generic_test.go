//go:build 1.18
// +build 1.18

package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpt_Get_Present(t *testing.T) {
	o := New(42)

	v, err := o.Get()
	assert.True(t, o.Present())
	assert.NoError(t, err)
	assert.Equal(t, 42, v)
}

func TestOpt_Get_NotPresent(t *testing.T) {
	o := Opt[int]{}
	var zero int

	v, err := o.Get()
	assert.False(t, o.Present())
	assert.Error(t, err)
	assert.Equal(t, zero, v)
}

func TestOpt_MustGet_Present(t *testing.T) {
	o := New(42)

	v := o.MustGet()
	assert.True(t, o.Present())
	assert.Equal(t, 42, v)
}

func TestOpt_MustGet_NotPresent(t *testing.T) {
	o := Opt[int]{}

	assert.Panics(t, func() { _ = o.MustGet() })
	assert.False(t, o.Present())
}

func TestOpt_OrElse_Present(t *testing.T) {
	o := New(42)

	v := o.OrElse(99)
	assert.True(t, o.Present())
	assert.Equal(t, 42, v)
}

func TestOpt_OrElse_NotPresent(t *testing.T) {
	o := Opt[int]{}

	v := o.OrElse(99)
	assert.False(t, o.Present())
	assert.Equal(t, 99, v)
}

func TestOpt_If_Present(t *testing.T) {
	o := New(42)

	canary := false
	o.If(func(v int) {
		canary = true
	})
	assert.True(t, o.Present())
	assert.True(t, canary)
}

func TestOpt_If_NotPresent(t *testing.T) {
	o := Opt[int]{}

	canary := false
	o.If(func(v int) {
		canary = true
	})
	assert.False(t, o.Present())
	assert.False(t, canary)
}

func TestOpt_MarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     Opt[int]
		WithZeroValue Opt[int]
		WithNoValue   Opt[int]
		Unused        Opt[int]
	}

	var instance = fields{
		WithValue:     New(42),
		WithZeroValue: New(0),
		WithNoValue:   Opt[int]{},
	}

	out, err := json.Marshal(instance)
	assert.NoError(t, err)
	assert.Equal(t, `{"WithValue":42,"WithZeroValue":0,"WithNoValue":null,"Unused":null}`, string(out))
}

func TestOpt_UnmarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     Opt[int]
		WithZeroValue Opt[int]
		WithNoValue   Opt[int]
		Unused        Opt[int]
	}

	var jsonString = `{"WithValue":42,"WithZeroValue":0,"WithNoValue":null}`
	instance := fields{}

	err := json.Unmarshal([]byte(jsonString), &instance)
	assert.NoError(t, err)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, 42, *instance.WithValue.value)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, 0, *instance.WithZeroValue.value)

	assert.False(t, instance.WithNoValue.Present())
	assert.Nil(t, instance.WithNoValue.value)

	assert.False(t, instance.Unused.Present())
	assert.Nil(t, instance.Unused.value)
}

func TestOpt_UnmarshalJSON_Overwritten(t *testing.T) {
	type fields struct {
		WithValue     Opt[int]
		WithZeroValue Opt[int]
		WithNoValue   Opt[int]
		Unused        Opt[int]
	}

	var jsonString = `{"WithValue":42,"WithZeroValue":0,"WithNoValue":null}`
	instance := fields{
		WithValue:     New(1),
		WithZeroValue: New(2),
		WithNoValue:   New(3),
		Unused:        New(4),
	}

	err := json.Unmarshal([]byte(jsonString), &instance)
	assert.NoError(t, err)

	assert.True(t, instance.WithValue.Present())
	assert.Equal(t, 42, *instance.WithValue.value)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, 0, *instance.WithZeroValue.value)

	assert.False(t, instance.WithNoValue.Present())
	assert.Nil(t, instance.WithNoValue.value)

	assert.True(t, instance.Unused.Present())
	assert.Equal(t, 4, *instance.Unused.value)
}
