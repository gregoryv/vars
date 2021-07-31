package vars

import (
	"testing"
)

func TestCopy_non_pair(t *testing.T) {
	defer catchPanic(t)
	var i int
	err := Copy(&i) // should panic
	if err != nil { // should not happen
		t.Fatal(err)
	}
}

func TestCopyAll(t *testing.T) {
	var i int
	var s string

	pairs := []interface{}{
		&i, 1, // ok
		&s, "a", // ok
		&i, "b", // bad
		&s, 0, // bad
	}

	if err := CopyAll(pairs[:2]...); err != nil {
		t.Fatal(err)
	}
	err := CopyAll(pairs...)
	if err == nil {
		t.Fatal("should fail")
	}
	if err.Len() != 2 {
		t.Error("wrong error len", err.Len())
	}
	if err.Error() == "" {
		t.Error("missing combined error")
	}
	if l := err.List(); len(l) != err.Len() {
		t.Error("List inconsistent with Len")
	}
}

func TestCopy(t *testing.T) {
	bad := struct{}{}

	var s string
	if err := Copy(&s, "john"); err != nil {
		t.Error(err)
	}
	badCopy(t, &s, bad)

	var r rune // same as int32
	if err := Copy(&r, 'k'); err != nil {
		t.Error(err)
	}
	badCopy(t, &r, bad)

	var i32 int32 // same as rune
	if err := Copy(&i32, 'k'); err != nil {
		t.Error(err)
	}
	badCopy(t, &i32, bad)

	var i int
	if err := Copy(&i, 54); err != nil {
		t.Error(err)
	}
	badCopy(t, &i, bad)

	var u uint
	if err := Copy(&u, uint(7)); err != nil {
		t.Error(err)
	}
	badCopy(t, &u, bad)

	var f float64
	if err := Copy(&f, 54.0); err != nil {
		t.Error(err)
	}
	badCopy(t, &f, bad)

	var b bool
	if err := Copy(&b, true); err != nil {
		t.Error(err)
	}
	badCopy(t, &b, bad)

	// bad destination
	badCopy(t, bad, bad)
}

func badCopy(t *testing.T, dst, src interface{}) {
	t.Helper()
	if err := Copy(dst, src); err == nil {
		t.Errorf("%T <- %T should fail", dst, src)
	}
}

func catchPanic(t *testing.T) {
	if recover() == nil {
		t.Helper()
		t.Error("should panic")
	}
}
