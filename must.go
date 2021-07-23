package vars

// MustCopy is the same as Copy only panics on error
func MustCopy(pairs ...interface{}) {
	if err := Copy(pairs...); err != nil {
		panic(err)
	}
}

// MustCopyAll is the same as CopyAll only panics on error
func MustCopyAll(pairs ...interface{}) {
	if err := CopyAll(pairs...); err != nil {
		panic(err)
	}
}
