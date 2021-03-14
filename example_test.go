package optional_test

import (
	"encoding/json"
	"fmt"

	"github.com/markphelps/optional"
)

func Example_get() {
	values := []optional.String{
		optional.NewString("foo"),
		optional.NewString(""),
		optional.NewString("bar"),
		{},
	}

	for _, v := range values {
		value, err := v.Get()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}

	// Output:
	// foo
	//
	// bar
	// value not present
}

func Example_mustGet() {
	values := []optional.String{
		optional.NewString("foo"),
		optional.NewString(""),
		optional.NewString("bar"),
		{},
	}

	for _, v := range values {
		if v.Present() {
			value := v.MustGet()
			fmt.Println(value)
		} else {
			fmt.Println("not present")
		}
	}

	// Output:
	// foo
	//
	// bar
	// not present
}

func Example_orElse() {
	values := []optional.String{
		optional.NewString("foo"),
		optional.NewString(""),
		optional.NewString("bar"),
		{},
	}

	for _, v := range values {
		fmt.Println(v.OrElse("not present"))
	}

	// Output:
	// foo
	//
	// bar
	// not present
}

func Example_if() {
	values := []optional.String{
		optional.NewString("foo"),
		optional.NewString(""),
		optional.NewString("bar"),
		{},
	}

	for _, v := range values {
		v.If(func(s string) {
			fmt.Println("present")
		})
	}

	// Output:
	// present
	// present
	// present
}

func Example_set() {
	var (
		values = []string{
			"foo",
			"",
			"bar",
		}

		s = optional.NewString("baz")
	)

	for _, v := range values {
		s.Set(v)
		value, err := s.Get()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}

	// Output:
	// foo
	//
	// bar
}

func Example_marshalJSON() {
	type example struct {
		Field    *optional.String `json:"field,omitempty"`
		FieldTwo *optional.String `json:"field_two"`
	}

	var values = []optional.String{
		optional.NewString("foo"),
		optional.NewString(""),
		optional.NewString("bar"),
	}

	for _, v := range values {
		out, _ := json.Marshal(&example{
			Field:    &v,
			FieldTwo: &v,
		})
		fmt.Println(string(out))
	}

	out, _ := json.Marshal(&example{})
	fmt.Println(string(out))

	// Output:
	// {"field":"foo","field_two":"foo"}
	// {"field":"","field_two":""}
	// {"field":"bar","field_two":"bar"}
	// {"field_two":null}
}

func Example_unmarshalJSON() {
	var values = []string{
		`{"field":"foo","field_two":"foo"}`,
		`{"field":"","field_two":""}`,
		`{"field":"null","field_two":"null"}`,
		`{"field":"bar","field_two":"bar"}`,
		"{}",
	}

	for _, v := range values {
		var o = &struct {
			Field    optional.String `json:"field,omitempty"`
			FieldTwo optional.String `json:"field_two"`
		}{}

		if err := json.Unmarshal([]byte(v), o); err != nil {
			fmt.Println(err)
		}

		o.Field.If(func(s string) {
			fmt.Println(s)
		})

		o.FieldTwo.If(func(s string) {
			fmt.Println(s)
		})
	}

	// Output:
	// foo
	// foo
	//

	//
	// bar
	// bar
}
