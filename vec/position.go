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
		if e, ok := v.na[val]; ok {
			nas[ix] = e
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
		if e, ok := v.na[val]; ok {
			nas[ix] = e
		}
	}
	return StrVector{
		obs:  new,
		na:   nas,
		size: len(new),
	}
}

// Find ...
func (v IntVector) Find(f func(v int) bool) []bool {
	new := make([]bool, v.Size())
	for ix, val := range v.obs {
		if _, ok := v.na[ix]; ok {
			new[ix] = false
		} else {
			new[ix] = f(val)
		}
	}
	return new
}

// Find ...
func (v StrVector) Find(f func(v string) bool) []bool {
	new := make([]bool, v.Size())

	for ix, val := range v.obs {
		if _, ok := v.na[ix]; ok {
			new[ix] = false
		} else {
			new[ix] = f(val)
		}
	}
	return new
}

// Filter ...
func (v IntVector) Filter(f func(v int) bool) Vector {
	locb := v.Find(f)
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
	locb := v.Find(f)
	new := make([]int, 0, v.Size())
	for ix, val := range locb {
		if val {
			new = append(new, ix)
		}
	}
	return v.Loc(new)
}
