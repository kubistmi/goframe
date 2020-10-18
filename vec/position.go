package vec

import "fmt"

// Loc ...
func (v IntVector) Loc(p []int) Vector {
	new := make([]int, len(p))
	nas := make(Set)
	for ix, val := range p {
		if val >= v.Size() {
			return IntVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
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
	nas := make(Set, len(p))
	for ix, val := range p {
		if val >= v.Size() {
			return StrVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
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

// Mask ...
func (v IntVector) Mask(f func(v int) bool) []bool {
	new := make([]bool, v.Size())
	for ix, val := range v.obs {
		if v.na.Get(ix) {
			new[ix] = false
		} else {
			new[ix] = f(val)
		}
	}
	return new
}

// Mask ...
func (v StrVector) Mask(f func(v string) bool) []bool {
	new := make([]bool, v.Size())

	for ix, val := range v.obs {
		if v.na.Get(ix) {
			new[ix] = false
		} else {
			new[ix] = f(val)
		}
	}
	return new
}

// Filter ...
func (v IntVector) Filter(f func(v int) bool) Vector {
	locb := v.Mask(f)
	new := make([]int, 0, v.Size())
	for ix, val := range locb {
		if val {
			new = append(new, ix)
		}
	}
	return v.Loc(new)
}

// Filter ...
func (v StrVector) Filter(f func(v string) bool) Vector {
	locb := v.Mask(f)
	new := make([]int, 0, v.Size())
	for ix, val := range locb {
		if val {
			new = append(new, ix)
		}
	}
	return v.Loc(new)
}
