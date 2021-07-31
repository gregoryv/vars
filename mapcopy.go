package vars

import (
	"fmt"
)

func MapCopyAll(in map[string]interface{}, pairs ...interface{}) error {
	mr := &mapCopier{
		Input:           in,
		AllowMissingKey: false,
	}
	return mr.Copy(pairs...)
}

func MapCopy(in map[string]interface{}, pairs ...interface{}) error {
	mr := &mapCopier{
		Input:           in,
		AllowMissingKey: true,
	}
	return mr.Copy(pairs...)
}

// mapCopier is a nexus for reading values from a map
type mapCopier struct {
	Input           map[string]interface{}
	missing         []string
	AllowMissingKey bool
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
func (me *mapCopier) Copy(pairs ...interface{}) error {
	var errs errors
	for i := 0; i < len(pairs); i = i + 2 {
		dst, key := pairs[i], pairs[i+1].(string)
		val, found := me.Input[key]
		if !found {
			me.missing = append(me.missing, key)
			continue
		}
		err := Copy(dst, val)
		if err != nil {
			errs.add(fmt.Errorf("%s %s", key, err.Error()))
		}
	}
	if !me.AllowMissingKey && len(me.missing) > 0 {
		for _, key := range me.missing {
			errs.add(fmt.Errorf("missing %q", key))
		}
	}
	if errs.len() != 0 {
		return &errs
	}
	return nil
}
