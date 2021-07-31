package vars

import "fmt"

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
	err := MapCopy(data,
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
	err := MapCopyAll(data,
		&addr, "addr",
		&name, "name",
	)
	if err != nil {
		fmt.Println(err)
	}
	// output:
	// MapCopyAll: missing addr, name
}
