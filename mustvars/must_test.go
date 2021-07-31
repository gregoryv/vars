package mustvars

import "testing"

func TestCopy(t *testing.T) {
	defer catchPanic(t)
	var s string
	Copy(&s, 0)
}

func TestCopyAll(t *testing.T) {
	defer catchPanic(t)
	var s string
	CopyAll(&s, "", &s, 1)
}

func catchPanic(t *testing.T) {
	if recover() == nil {
		t.Helper()
		t.Error("should panic")
	}
}
