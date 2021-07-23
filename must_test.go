package vars

import "testing"

func TestMustCopy(t *testing.T) {
	defer catchPanic(t)
	var s string
	MustCopy(&s, 0)
}

func TestMustCopyAll(t *testing.T) {
	defer catchPanic(t)
	var s string
	MustCopyAll(&s, "", &s, 1)
}

func catchPanic(t *testing.T) {
	if recover() == nil {
		t.Helper()
		t.Error("should panic")
	}
}
