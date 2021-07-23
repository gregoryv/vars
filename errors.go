package vars

import "strings"

type Errors struct {
	errors []error
}

func (me *Errors) add(err error) {
	me.errors = append(me.errors, err)
}

// Len returns number of errors
func (me *Errors) Len() int {
	return len(me.errors)
}

// Error returns a comma separated string of all the errors
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

// List returns a slice of all errors
func (me *Errors) List() []error {
	return me.errors
}
