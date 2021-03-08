package vec

// Err  ...
func (v IntVector) Err() error {
	return v.err
}

// Err  ...
func (v StrVector) Err() error {
	return v.err
}

func (v IntVector) setError(err error) Vector {
	v.err = err
	return v
}

func (v StrVector) setError(err error) Vector {
	v.err = err
	return v
}

func (v IntVector) setStrError(err error) StrVector {
	var e StrVector
	e.err = err
	return e
}

func (v StrVector) setStrError(err error) StrVector {
	v.err = err
	return v
}

func (v IntVector) setIntError(err error) IntVector {
	v.err = err
	return v
}

func (v StrVector) setIntError(err error) IntVector {
	var e IntVector
	e.err = err
	return e
}
