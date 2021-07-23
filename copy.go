// Package vars provides copy funcs for variables
package vars

import (
	"fmt"
	"strings"
)

// MustCopy is the same as Copy only panics on error
func MustCopy(pairs ...interface{}) {
	if err := Copy(pairs...); err != nil {
		panic(err)
	}
}

// Copy copies pairwise destination, source. Returns error on first
// error. Panics if pairs is uneven.
func Copy(pairs ...interface{}) error {
	for i := 0; i < len(pairs); i = i + 2 {
		dst, src := pairs[i], pairs[i+1]
		err := copyX(dst, src)
		if err != nil {
			return err
		}
	}
	return nil
}

func CopyAll(pairs ...interface{}) *Errors {
	result := &Errors{}
	for i := 0; i < len(pairs); i = i + 2 {
		dst, src := pairs[i], pairs[i+1]
		err := copyX(dst, src)
		if err != nil {
			result.add(err)
		}
	}
	if result.Len() > 0 {
		return result
	}
	return nil
}

type Errors struct {
	errors []error
}

func (me *Errors) add(err error) {
	me.errors = append(me.errors, err)
}

func (me *Errors) Len() int {
	return len(me.errors)
}

func (me *Errors) Error() string {
	var sb strings.Builder
	for i, err := range me.errors {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}

func (me *Errors) List() []error {
	return me.errors
}

func copyX(dst, src interface{}) error {
	switch dst := dst.(type) {
	case *string:
		return copyString(dst, src)
	case *int:
		return copyInt(dst, src)

	case *int32:
		return copyInt32(dst, src)

	case *uint:
		return copyUInt(dst, src)
	case *float64:
		return copyFloat64(dst, src)
	case *bool:
		return copyBool(dst, src)

	default:
		return fmt.Errorf("cannot copy %T to %T", src, dst)
	}
}

func copyString(dst *string, src interface{}) error {
	v, ok := src.(string)
	if !ok {
		return fmt.Errorf("not string")
	}
	*dst = v
	return nil
}

func copyInt(dst *int, src interface{}) error {
	v, ok := src.(int)
	if !ok {
		return fmt.Errorf("not int")
	}
	*dst = v
	return nil
}

func copyInt32(dst *int32, src interface{}) error {
	v, ok := src.(int32)
	if !ok {
		return fmt.Errorf("not int32")
	}
	*dst = v
	return nil
}

func copyUInt(dst *uint, src interface{}) error {
	v, ok := src.(uint)
	if !ok {
		return fmt.Errorf("not uint")
	}
	*dst = v
	return nil
}

func copyFloat64(dst *float64, src interface{}) error {
	v, ok := src.(float64)
	if !ok {
		return fmt.Errorf("not float64")
	}
	*dst = v
	return nil
}

func copyBool(dst *bool, src interface{}) error {
	v, ok := src.(bool)
	if !ok {
		return fmt.Errorf("not bool")
	}
	*dst = v
	return nil
}
