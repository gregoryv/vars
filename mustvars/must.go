// Package mustvars provides wrappers which panic on error.
package mustvars

import (
	"github.com/gregoryv/vars"
)

// Copy wraps vars.Copy only panics on error
func Copy(pairs ...interface{}) {
	if err := vars.Copy(pairs...); err != nil {
		panic(err)
	}
}

// CopyAll wraps vars.CopyAll only panics on error
func CopyAll(pairs ...interface{}) {
	if err := vars.CopyAll(pairs...); err != nil {
		panic(err)
	}
}
