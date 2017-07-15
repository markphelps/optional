package optional_test

import (
	"testing"

	"github.com/markphelps/optional"
	"github.com/stretchr/testify/assert"
)

func TestOfString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want optional.String
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: optional.OfString(""),
		},
		{
			name: "non-empty string",
			args: args{s: "foo"},
			want: optional.OfString("foo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optional.OfString(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSet(t *testing.T) {
	type fields struct {
		s string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "empty string",
			fields: fields{},
			args:   args{s: ""},
		},
		{
			name:   "non-empty string",
			fields: fields{},
			args:   args{s: "foo"},
		},
		{
			name:   "pre-existing string",
			fields: fields{s: "foo"},
			args:   args{s: "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := optional.OfString(tt.fields.s)
			o.Set(tt.args.s)
			assert.Equal(t, o.String(), tt.args.s)
			assert.True(t, o.Present())
		})
	}
}

func TestPresent(t *testing.T) {
	type fields struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
		zeroed bool
	}{
		{
			name:   "non-existent",
			want:   false,
			zeroed: true,
		},
		{
			name:   "empty",
			fields: fields{},
			want:   true,
		},
		{
			name:   "non-empty",
			fields: fields{s: "foo"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var o optional.String
			if !tt.zeroed {
				o = optional.OfString(tt.fields.s)
			}
			assert.Equal(t, tt.want, o.Present())
		})
	}
}

func TestOrElse(t *testing.T) {
	type fields struct {
		s string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		zeroed bool
	}{
		{
			name:   "non-existent",
			args:   args{s: "foo"},
			want:   "foo",
			zeroed: true,
		},
		{
			name:   "empty",
			fields: fields{},
			args:   args{s: "bar"},
			want:   "",
		},
		{
			name:   "non-empty",
			fields: fields{s: "foo"},
			args:   args{s: "bar"},
			want:   "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var o optional.String
			if !tt.zeroed {
				o = optional.OfString(tt.fields.s)
			}
			got := o.OrElse(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
