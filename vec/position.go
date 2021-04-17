package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// IntVector implementations ---------------------------------------------------

// Loc prepares a new Data with the selected positions
func (v IntVector) Loc(p []int) Vector {
	new := make([]int, len(p))
	na := NewNA(-1)
	for ix, val := range p {
		if val >= v.Size() || val < 0 {
			return NewErrVec(fmt.Errorf("%w wrong position `p`, maximum allowed: %v, got %v", utils.ErrParamVal, v.Size()-1, val), v.Type())
		}
		new[ix] = v.data[val]
		na.Setif(v.na.Get(val), ix)
	}
	return unsafeIntVec(new, na)
}

// Check ...
func (v IntVector) Check(f interface{}) ([]bool, error) {
	fun, ok := f.(func(int, bool) bool)
	if !ok {
		if err, ok := f.(error); ok {
			return nil, fmt.Errorf("interface `f` contains error: %w", err)
		}
		return nil, fmt.Errorf("%w parameter f, expected: `func(int) bool`, got `%T`", utils.ErrParamType, f)
	}
	new := make([]bool, v.Size())
	for ix, val := range v.data {
		new[ix] = fun(val, v.na.Get(ix))
	}
	return new, nil
}

// Filter ...
func (v IntVector) Filter(f interface{}) Vector {
	locb, err := v.Check(f)
	if err != nil {
		return NewErrVec(fmt.Errorf("error in Check: %w", err), v.Type())
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
		return NewErrVec(fmt.Errorf("%w size of slice `c` != size of Vector `v`, expected: %v, got: %v", utils.ErrParamVal, v.Size(), len(c)), v.Type())
	}
	pos := make([]int, 0, v.Size())
	for ix, val := range c {
		if val {
			pos = append(pos, ix)
		}
	}
	return v.Loc(pos)
}

// StrVector implementations ---------------------------------------------------

// Loc prepares a new Data with the selected positions
func (v StrVector) Loc(p []int) Vector {
	new := make([]string, len(p))
	na := make(Set)
	for ix, val := range p {
		if val >= v.Size() || val < 0 {
			return NewErrVec(fmt.Errorf("%w wrong position `p`, maximum allowed: %v, got %v", utils.ErrParamVal, v.Size()-1, val), v.Type())
		}
		new[ix] = v.data[val]
		na.Setif(v.na.Get(val), ix)
	}
	return unsafeStrVec(new, na)
}

// Check ...
func (v StrVector) Check(f interface{}) ([]bool, error) {
	fun, ok := f.(func(string, bool) bool)
	if !ok {
		if err, ok := f.(error); ok {
			return nil, fmt.Errorf("interface `f` contains error: %w", err)
		}
		return nil, fmt.Errorf("%w parameter f, expected: `func(string) bool`, got `%T`", utils.ErrParamType, f)
	}

	new := make([]bool, v.Size())
	for ix, val := range v.data {
		new[ix] = fun(val, v.na.Get(ix))
	}
	return new, nil
}

// Filter ...
func (v StrVector) Filter(f interface{}) Vector {
	locb, err := v.Check(f)
	if err != nil {
		return NewErrVec(fmt.Errorf("error in Check: %w", err), v.Type())
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
func (v StrVector) Mask(c []bool) Vector {
	if len(c) != v.Size() {
		return NewErrVec(fmt.Errorf("%w size of slice `c` != size of Vector `v`, expected: %v, got: %v", utils.ErrParamVal, v.Size(), len(c)), v.Type())
	}
	pos := make([]int, 0, v.Size())
	for ix, val := range c {
		if val {
			pos = append(pos, ix)
		}
	}
	return v.Loc(pos)
}

// BoolVector implementations ---------------------------------------------------

// Loc prepares a new Data with the selected positions
func (v BoolVector) Loc(p []int) Vector {
	new := make([]bool, len(p))
	na := NewNA(-1)
	for ix, val := range p {
		if val >= v.Size() || val < 0 {
			return NewErrVec(fmt.Errorf("%w wrong position `p`, maximum allowed: %v, got %v", utils.ErrParamVal, v.Size()-1, val), v.Type())
		}
		new[ix] = v.data[val]
		na.Setif(v.na.Get(val), ix)
	}
	return unsafeBoolVec(new, na)
}

// Check ...
func (v BoolVector) Check(f interface{}) ([]bool, error) {
	fun, ok := f.(func(bool, bool) bool)
	if !ok {
		if err, ok := f.(error); ok {
			return nil, fmt.Errorf("interface `f` contains error: %w", err)
		}
		return nil, fmt.Errorf("%w parameter f, expected: `func(int) bool`, got `%T`", utils.ErrParamType, f)
	}
	new := make([]bool, v.Size())
	for ix, val := range v.data {
		new[ix] = fun(val, v.na.Get(ix))
	}
	return new, nil
}

// Filter ...
func (v BoolVector) Filter(f interface{}) Vector {
	locb, err := v.Check(f)
	if err != nil {
		return NewErrVec(fmt.Errorf("error in Check: %w", err), v.Type())
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
func (v BoolVector) Mask(c []bool) Vector {
	if len(c) != v.Size() {
		return NewErrVec(fmt.Errorf("%w size of slice `c` != size of Vector `v`, expected: %v, got: %v", utils.ErrParamVal, v.Size(), len(c)), v.Type())
	}
	pos := make([]int, 0, v.Size())
	for ix, val := range c {
		if val {
			pos = append(pos, ix)
		}
	}
	return v.Loc(pos)
}
