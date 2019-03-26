package must

// Must panic if err not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Must2 panic if second arg err not nil.
func Must2(_ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}

// Must3 panic if third arg err not nil.
func Must3(_, _ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}

// Must2R panic if second arg err not nil and return first arg.
func Must2R(r interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}

	return r
}

// Must3R panic if third arg err not nil and return first two args.
func Must3R(r1, r2 interface{}, err error) (interface{}, interface{}) {
	if err != nil {
		panic(err)
	}

	return r1, r2
}
