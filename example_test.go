package optional_test

import (
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
		if v.Present() {
			value, err := v.Get()
			if err == nil {
				fmt.Println(value)
			}
		}
	}

	// Output:
	// foo
	//
	// bar
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
