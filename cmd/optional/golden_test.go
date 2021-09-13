package main

import (
	"bufio"
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var updateFlag = flag.Bool("update", true, "Update the golden files.")

type testcase struct {
	name        string
	packageName string
	outputName  string
	result      bool
	typeName    string
	filename    string
}

var tests = []testcase{
	{"built-in datatype", "testdata", "String", false, "string", "string.go"},
	{"custom exported type", "testdata", "OptionalFoo", false, "Foo", "optional_foo.go"},
	{"custom non-exported type", "testdata", "optionalBar", false, "bar", "optional_bar.go"},
	// result
	{"built-in datatype result", "testdata", "ResultString", true, "string", "result_string.go"},
	{"custom exported type result", "testdata", "ResultFoo", true, "Foo", "result_foo.go"},
	{"custom non-exported type result", "testdata", "resultBar", true, "bar", "result_bar.go"},
}

// TestGolden compares generated files with 'golden files' line by line
// to assert that generated files match what is expected
func TestGolden(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := generator{
				packageName: test.packageName,
				outputName:  test.outputName,
				typeName:    test.typeName,
				result:      test.result,
			}

			output, err := g.generate()
			if err != nil {
				t.Error(err)
			}

			if *updateFlag {
				err = ioutil.WriteFile("./testdata/"+test.filename, output, 0644)
				if err != nil {
					log.Fatalf("writing output: %s", err)
				}
				log.Printf("updated goldenfile: %s", "./testdata/"+test.filename)
			}

			input, err := ioutil.ReadFile("./testdata/" + test.filename)
			if err != nil {
				t.Error(err)
			}

			want := bufio.NewScanner(bytes.NewReader(input))
			got := bufio.NewScanner(bytes.NewReader(output))

			count := 0

			for want.Scan() {
				got.Scan()

				// skip comment lines at top
				if count < 2 {
					count++
					continue
				}

				assert.Equal(t, want.Bytes(), got.Bytes())
				count++
			}

			if err := want.Err(); err != nil {
				t.Error(err)
			}
			if err := got.Err(); err != nil {
				t.Error(err)
			}
		})
	}
}
