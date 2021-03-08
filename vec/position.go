package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// Loc ...
func (v IntVector) Loc(p []int) Vector {
	new := make([]int, len(p))
	nas := make(Set)
	for ix, val := range p {
		if val >= v.Size() || val < 0 {
			return v.setError(fmt.Errorf("%w wrong position `p`, maximum allowed: %v, got %v", utils.ErrParamVal, v.Size()-1, val))
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
			return v.setError(fmt.Errorf("%w wrong position `p`, maximum allowed: %v, got %v", utils.ErrParamVal, v.Size()-1, val))
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
	fun, ok := f.(func(int, bool) bool)
	if !ok {
		if err, ok := f.(error); ok {
			return nil, fmt.Errorf("interface `f` contains error: %w", err)
		}
		return nil, fmt.Errorf("%w parameter f, expected: `func(int) bool`, got `%T`", utils.ErrParamType, f)
	}
	new := make([]bool, v.Size())
	for ix, val := range v.obs {
		new[ix] = fun(val, v.na.Get(ix))
	}
	return new, nil
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
	for ix, val := range v.obs {
		new[ix] = fun(val, v.na.Get(ix))
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
		return v.setError(fmt.Errorf("%w size of slice `c` != size of Vector `v`, expected: %v, got: %v", utils.ErrParamVal, v.Size(), len(c)))
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
		return v.setError(fmt.Errorf("%w size of slice `c` != size of Vector `v`, expected: %v, got: %v", utils.ErrParamVal, v.Size(), len(c)))
	}
	pos := make([]int, 0, v.Size())
	for ix, val := range c {
		if val {
			pos = append(pos, ix)
		}
	}
	return v.Loc(pos)
}
