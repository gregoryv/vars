package vars

import "strings"

// SplitErr takes an error produces by CopyAll or MapCopyAll and
// splits it into multiple errors. Panics if in error is of other
// origin.
func SplitErr(in error) []error {
	return in.(errors)
}

type errors []error

// Error returns a comma separated string of all the errors
func (me errors) Error() string {
	var sb strings.Builder
	for i, err := range me {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}
