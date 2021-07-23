package vars_test

import (
	"fmt"

	"github.com/gregoryv/vars"
)

func ExampleCopy() {
	var (
		i int
		s string
	)
	pairs := []interface{}{
		&i, 1, // ok
		&s, "a", // ok
		&i, "b", // bad
		&s, 0, // bad
	}
	err := vars.CopyAll(pairs...)
	for _, err := range err.List() {
		fmt.Println(err)
	}
	// output:
	// Copy[5]: not int
	// Copy[7]: not string
}
