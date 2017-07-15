package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: String{string: "", present: true},
		},
		{
			name: "non-empty string",
			args: args{s: "foo"},
			want: String{string: "foo", present: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfString(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSet(t *testing.T) {
	type fields struct {
		string  string
		present bool
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
			fields: fields{string: "foo", present: true},
			args:   args{s: "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &String{
				string:  tt.fields.string,
				present: tt.fields.present,
			}
			o.Set(tt.args.s)
			assert.Equal(t, o.String(), tt.args.s)
			assert.True(t, o.Present())
		})
	}
}

func TestPresent(t *testing.T) {
	type fields struct {
		string  string
		present bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   false,
		},
		{
			name:   "non-empty",
			fields: fields{string: "foo", present: true},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &String{
				string:  tt.fields.string,
				present: tt.fields.present,
			}
			assert.Equal(t, tt.want, o.Present())
		})
	}
}

func TestOrElse(t *testing.T) {
	type fields struct {
		string  string
		present bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			args:   args{s: "bar"},
			want:   "bar",
		},
		{
			name:   "non-empty",
			fields: fields{string: "foo", present: true},
			args:   args{s: "bar"},
			want:   "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &String{
				string:  tt.fields.string,
				present: tt.fields.present,
			}
			got := o.OrElse(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
