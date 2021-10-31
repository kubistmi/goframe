package vec

// IntVector implementations ---------------------------------------------------

// Get returns a copy of the underlying data
func (v IntVector) GetI() (interface{}, NA, error) {
	if err := v.Err(); err != nil {
		return []int{}, nil, err
	}
	new := make([]int, v.Size())
	copy(new, v.data)
	newna := v.na.CopyNA()
	return new, newna, nil
}

// Get ...
func (v IntVector) Get() ([]int, NA, error) {
	if err := v.Err(); err != nil {
		return []int{}, nil, err
	}
	new := make([]int, v.Size())
	copy(new, v.data)
	return new, v.na.CopyNA(), nil
}

// Copy returns a copy of Vector
func (v IntVector) Copy() Vector {
	if err := v.Err(); err != nil {
		return NewErrVec(err, v.Type())
	}
	new := make([]int, v.Size())
	copy(new, v.data)
	newna := v.na.CopyNA()

	return IntVector{
		data: new,
		na:   newna,
	}
}

// StrVector implementations ---------------------------------------------------

// Get returns a copy of the underlying data
func (v StrVector) GetI() (interface{}, NA, error) {
	if err := v.Err(); err != nil {
		return []string{}, nil, err
	}
	new := make([]string, v.Size())
	copy(new, v.data)
	newna := v.na.CopyNA()
	return new, newna, nil
}

// Get ...
func (v StrVector) Get() ([]string, NA, error) {
	if err := v.Err(); err != nil {
		return []string{}, nil, err
	}
	new := make([]string, v.Size())
	copy(new, v.data)
	return new, v.na.CopyNA(), nil
}

// Copy returns a copy of Vector
func (v StrVector) Copy() Vector {
	if err := v.Err(); err != nil {
		return NewErrVec(err, v.Type())
	}
	new := make([]string, v.Size())
	copy(new, v.data)
	newna := v.na.CopyNA()

	return StrVector{
		data: new,
		na:   newna,
	}
}

func (v StrVector) Unique() []string {
	if v.index == nil {
		v = v.Group().Str()
	}

	out := make([]string, 0, len(v.index))
	for ix := range v.index {
		out = append(out, ix)
	}
	return out
}

// BoolVector implementations ---------------------------------------------------

// Get returns a copy of the underlying data
func (v BoolVector) GetI() (interface{}, NA, error) {
	if err := v.Err(); err != nil {
		return []bool{}, nil, err
	}
	new := make([]bool, v.Size())
	copy(new, v.data)
	newna := v.na.CopyNA()
	return new, newna, nil
}

// Get ...
func (v BoolVector) Get() ([]bool, NA, error) {
	if err := v.Err(); err != nil {
		return []bool{}, nil, err
	}
	new := make([]bool, v.Size())
	copy(new, v.data)
	return new, v.na.CopyNA(), nil
}

// Copy returns a copy of Vector
func (v BoolVector) Copy() Vector {
	if err := v.Err(); err != nil {
		return NewErrVec(err, v.Type())
	}
	new := make([]bool, v.Size())
	copy(new, v.data)
	newna := v.na.CopyNA()

	return BoolVector{
		data: new,
		na:   newna,
	}
}
