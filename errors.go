package vars

import "strings"

// SplitErr takes an error produced by CopyAll or MapCopyAll and
// splits it into multiple errors. Panics if in error is of other
// origin or nil.
func SplitErr(in error) []error {
	return in.(*errors).list
}

type errors struct {
	list []error
}

func (me *errors) add(e error) {
	me.list = append(me.list, e)
}

func (me *errors) len() int {
	return len(me.list)
}

// Error returns a comma separated string of all the errors
func (me *errors) Error() string {
	var sb strings.Builder
	for i, err := range me.list {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}
