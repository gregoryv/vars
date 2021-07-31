package vars_test

import (
	"fmt"

	"github.com/gregoryv/vars"
)

func ExampleCopyAll() {
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

	errs := vars.SplitErr(err)
	for _, err := range errs {
		fmt.Println(err)
	}
	// output:
	// Copy[5]: not int
	// Copy[7]: not string
}

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
	err := vars.Copy(pairs...)
	fmt.Println(err)
	// output:
	// Copy[5]: not int
}

func ExampleMapCopy() {
	data := map[string]interface{}{
		"name":   "John Doe",
		"age":    54,
		"weight": 94.5, // kilos
		"alive":  true,
	}

	var (
		name   string
		age    int
		weight float64
		alive  bool
		addr   string
	)
	err := vars.MapCopy(data,
		&name, "name",
		&age, "age",
		&weight, "weight",
		&alive, "alive",
		&addr, "addr", // missing values are ignored
	)
	if err != nil {
		fmt.Println(err)
	}
	// output:
}

func ExampleMapCopyAll() {
	data := map[string]interface{}{}

	var addr, name string
	err := vars.MapCopyAll(data,
		&addr, "addr",
		&name, "name",
	)
	if err != nil {
		fmt.Println(err)
		// or each error
		for _, err := range vars.SplitErr(err) {
			fmt.Println(err)
		}
	}
	// output:
	// missing "addr", missing "name"
	// missing "addr"
	// missing "name"
}
