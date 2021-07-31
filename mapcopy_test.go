package vars

import (
	"strings"
	"testing"
)

var data = map[string]interface{}{
	"name":   "John Doe",
	"age":    54,
	"weight": 94.5, // kilos
	"alive":  true,
}

func TestMapCopy_all_fields_exist(t *testing.T) {
	var (
		name   string
		age    int
		weight float64
		alive  bool
	)
	err := MapCopy(data,
		&name, "name",
		&age, "age",
		&weight, "weight",
		&alive, "alive",
	)
	if err != nil {
		t.Error(err)
	}
}

func TestMapCopy_ignore_missing(t *testing.T) {
	var (
		name string
		addr string
	)
	err := MapCopy(data,
		&name, "name",
		&addr, "addr", // missing values are ignored
	)
	if err != nil {
		t.Error(err)
	}
}

func TestMapCopy_fails_on_type_missmatch(t *testing.T) {
	var (
		name string
	)
	err := MapCopy(data,
		&name, "age", // string <- int
	)
	if err == nil {
		t.Fail()
	}
	if !strings.Contains(err.Error(), "age") {
		t.Errorf("missing %q in error: %s", "age", err)
	}
}

func TestMapCopy_skips_subsequent_errors(t *testing.T) {
	var (
		name string
	)
	err := MapCopy(data,
		&name, "age", // string <- int
		&name, "weight", // second is skipped
	)
	if err == nil {
		t.Fail()
	}
	t.Log("got", err)

	if !strings.Contains(err.Error(), "weight") {
		t.Errorf("missing %q", "weight")
	}
	if !strings.Contains(err.Error(), "age") {
		t.Errorf("missing %q", "age")
	}
}

func TestMapCopyAll_fail_with_all_errors(t *testing.T) {
	var (
		name string
		x    int
	)
	err := MapCopyAll(data,
		&name, "name", // ok
		&name, "addr", // missing
		&x, "name", // bad type
	)
	if err == nil {
		t.Fatal("worked even when missing keys and bad types")
	}
	errs := SplitErr(err)
	if len(errs) != 2 {
		t.Error("got", err)
	}
}
