package vars

import (
	"fmt"
)

func MapCopyAll(in map[string]interface{}, pairs ...interface{}) error {
	mr := &mapCopier{Input: in}
	mr.Copy(pairs...)
	err := mr.Error()
	if err != nil {
		return err
	}
	if len(mr.missing) > 0 {
		errs := make(errors, len(mr.missing))
		for i, key := range mr.missing {
			errs[i] = fmt.Errorf("missing %q", key)
		}
		return errs
	}
	return nil
}

func MapCopy(in map[string]interface{}, pairs ...interface{}) error {
	mr := &mapCopier{Input: in}
	mr.Copy(pairs...)
	return mr.Error()
}

// mapCopier is a nexus for reading values from a map
type mapCopier struct {
	Input   map[string]interface{}
	missing []string
	err     error
}

// Read copies pairs of values from the input map.
// Pairs argument must be an even list of destination followed by key in map.
//
//  [
//    dest, key,
//    dest, key,
//    ...,
//    dest, key,
//  ]
//
func (me *mapCopier) Copy(pairs ...interface{}) {
	for i := 0; i < len(pairs); i = i + 2 {
		if me.err != nil {
			return
		}
		dst, key := pairs[i], pairs[i+1].(string)
		val, found := me.Input[key]
		if !found {
			me.missing = append(me.missing, key)
			continue
		}
		me.err = prefixErr(key, Copy(dst, val))
	}
}

func (me *mapCopier) Error() error {
	return me.err
}

func prefixErr(name string, err error) error {
	if err != nil {
		return fmt.Errorf("%s %s", name, err.Error())
	}
	return nil
}
