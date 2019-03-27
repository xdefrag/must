package must

type errorer interface {
	Error(args ...interface{})
}

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

// TMust takes first parameters errorer interface and returns function
// which call Error(err) on errorer if first parameter err not nil.
func TMust(t errorer) func(error) {
	return func(err error) {
		if err != nil {
			t.Error(err)
		}
	}
}

// TMust2 takes first parameters errorer interface and returns function
// which call Error(err) on errorer if second parameter err not nil.
func TMust2(t errorer) func(interface{}, error) {
	return func(_ interface{}, err error) {
		if err != nil {
			t.Error(err)
		}
	}
}

// TMust3 takes first parameters errorer interface and returns function
// which call Error(err) on errorer if third parameter err not nil.
func TMust3(t errorer) func(interface{}, interface{}, error) {
	return func(_, _ interface{}, err error) {
		if err != nil {
			t.Error(err)
		}
	}
}
