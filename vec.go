package main

import "fmt"

// Vector ...
type Vector interface {
	Size() int
	//Get[T vector_type]() []T
	GetI() (interface{}, map[int]bool)
	Loc(p []int) Vector
	Err() error
	Copy() Vector
}

// NewVec ...
func NewVec(data interface{}, na ...map[int]bool) Vector {
	switch t := data.(type) {
	case []int:
		new := make([]int, len(t))
		copy(new, t)
		return IntVector{
			obs:  new,
			na:   na[0],
			size: len(t),
		}

	case []string:
		new := make([]string, len(t))
		copy(new, t)
		return StrVector{
			obs:  new,
			na:   na[0],
			size: len(t),
		}
	default:
		return StrVector{
			err: fmt.Errorf("wrong data type, expected []int or []string, got %T", t),
		}
	}
}

func cpMap(m map[int]bool) map[int]bool {
	new := make(map[int]bool)
	for key, val := range m {
		new[key] = val
	}
	return new
}

// IntVector ... ---------------------------------------------------------------
type IntVector struct {
	obs   []int
	na    map[int]bool
	index []int
	size  int
	err   error
}

// Size ...
func (v IntVector) Size() int {
	return v.size
}

// GetI ...
func (v IntVector) GetI() (interface{}, map[int]bool) {
	new := make([]int, v.Size())
	copy(new, v.obs)
	nas := cpMap(v.na)
	return new, nas
}

// Get ...
func (v IntVector) Get() ([]int, map[int]bool) {
	new := make([]int, v.Size())
	copy(new, v.obs)
	nas := cpMap(v.na)
	return new, nas
}

// Err  ...
func (v IntVector) Err() error {
	return v.err
}

// Loc ...
func (v IntVector) Loc(p []int) Vector {
	new := make([]int, len(p))
	nas := make(map[int]bool)
	for ix, val := range p {
		if val >= v.Size() {
			return IntVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
		}
		new[ix] = v.obs[val]
		if _, ok := v.na[val]; ok {
			nas[ix] = true
		}
	}
	return IntVector{
		obs:  new,
		na:   nas,
		size: len(p),
	}
}

// Copy ...
func (v IntVector) Copy() Vector {

	new := make([]int, v.size)
	copy(new, v.obs)
	nas := cpMap(v.na)

	return IntVector{
		obs:   new,
		na:    nas,
		size:  v.size,
		index: v.index,
		err:   v.err,
	}
}

// Mutate ...
func (v IntVector) Mutate(f func(v int) int) Vector {
	new := make([]int, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
	}
	return IntVector{
		obs:   new,
		na:    v.na,
		index: v.index,
		size:  v.size,
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

// StrVector ... ---------------------------------------------------------------
type StrVector struct {
	obs     []string
	na      map[int]bool
	index   []int
	size    int
	inverse map[string][]int //?inverse index
	err     error
}

// Size ...
func (v StrVector) Size() int {
	return v.size
}

// GetI ...
func (v StrVector) GetI() (interface{}, map[int]bool) {
	new := make([]string, v.Size())
	copy(new, v.obs)
	nas := cpMap(v.na)
	return new, nas
}

// Get ...
func (v StrVector) Get() ([]string, map[int]bool) {
	new := make([]string, v.Size())
	copy(new, v.obs)
	nas := cpMap(v.na)
	return new, nas
}

// Err  ...
func (v StrVector) Err() error {
	return v.err
}

// Loc ...
func (v StrVector) Loc(p []int) Vector {

	new := make([]string, len(p))
	nas := make(map[int]bool)
	for ix, val := range p {
		if val >= v.Size() {
			return StrVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
		}
		new[ix] = v.obs[val]
		if _, ok := v.na[val]; ok {
			nas[ix] = true
		}
	}
	return StrVector{
		obs:  new,
		na:   nas,
		size: len(new),
	}
}

// Copy ...
func (v StrVector) Copy() Vector {

	new := make([]string, v.size)
	copy(new, v.obs)
	nas := cpMap(v.na)

	return StrVector{
		obs:   new,
		na:    nas,
		size:  v.size,
		index: v.index,
		err:   v.err,
	}
}

// Mutate ...
func (v StrVector) Mutate(f func(v string) string) Vector {
	new := make([]string, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
	}
	return StrVector{
		obs:   new,
		na:    v.na,
		index: v.index,
		size:  v.size,
	}
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
