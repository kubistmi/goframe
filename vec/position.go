package vec

import "fmt"

// Loc ...
func (v IntVector) Loc(p []int) Vector {
	new := make([]int, len(p))
	nas := make(Set)
	for ix, val := range p {
		if val >= v.Size() || val < 0 {
			return NewErrVec(fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val))
		}
		new[ix] = v.obs[val]
		if v.na.Get(val) {
			nas = nas.Set(val)
		}
	}
	return IntVector{
		obs:  new,
		na:   nas,
		size: len(p),
	}
}

// Loc ...
func (v StrVector) Loc(p []int) Vector {
	new := make([]string, len(p))
	nas := make(Set)
	for ix, val := range p {
		if val >= v.Size() || val < 0 {
			return NewErrVec(fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val))
		}
		new[ix] = v.obs[val]
		if v.na.Get(val) {
			nas = nas.Set(val)
		}
	}
	return StrVector{
		obs:  new,
		na:   nas,
		size: len(new),
	}
}

// Check ...
func (v IntVector) Check(f interface{}) ([]bool, error) {
	fun, ok := f.(func(int) bool)
	if !ok {
		return nil, fmt.Errorf("wrong function, expected: `func(int) bool`, got `%T`", f)
	}

	new := make([]bool, v.Size())
	for ix, val := range v.obs {
		if v.na.Get(ix) {
			new[ix] = false
		} else {
			new[ix] = fun(val)
		}
	}
	return new, nil
}

// Check ...
func (v StrVector) Check(f interface{}) ([]bool, error) {
	fun, ok := f.(func(string) bool)
	if !ok {
		return nil, fmt.Errorf("wrong function, expected: `func(string) bool`, got `%T`", f)
	}

	new := make([]bool, v.Size())
	for ix, val := range v.obs {
		if v.na.Get(ix) {
			new[ix] = false
		} else {
			new[ix] = fun(val)
		}
	}
	return new, nil
}

// Filter ...
func (v IntVector) Filter(f interface{}) Vector {
	locb, err := v.Check(f)
	if err != nil {
		return NewErrVec(fmt.Errorf("error in Check: %w", err))
	}
	new := make([]int, 0, v.Size())
	for ix, val := range locb {
		if val {
			new = append(new, ix)
		}
	}
	return v.Loc(new)
}

// Filter ...
func (v StrVector) Filter(f interface{}) Vector {
	locb, err := v.Check(f)
	if err != nil {
		return NewErrVec(fmt.Errorf("error in Check: %w", err))
	}
	new := make([]int, 0, v.Size())
	for ix, val := range locb {
		if val {
			new = append(new, ix)
		}
	}
	return v.Loc(new)
}

// Mask ...
func (v IntVector) Mask(c []bool) Vector {
	if len(c) != v.Size() {
		return NewErrVec(fmt.Errorf("size of boolean slice does not match the size of Vector, expected: %v, got: %v", v.Size(), len(c)))
	}
	pos := make([]int, 0, v.Size())
	for ix, val := range c {
		if val {
			pos = append(pos, ix)
		}
	}
	return v.Loc(pos)
}

// Mask ...
func (v StrVector) Mask(c []bool) Vector {
	if len(c) != v.Size() {
		return NewErrVec(fmt.Errorf("size of boolean slice does not match the size of Vector, expected: %v, got: %v", v.Size(), len(c)))
	}
	pos := make([]int, 0, v.Size())
	for ix, val := range c {
		if val {
			pos = append(pos, ix)
		}
	}
	return v.Loc(pos)
}
