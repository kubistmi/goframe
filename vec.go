package main

import "fmt"

// Vector ...
type Vector interface {
	Size() int
	//Get[T vector_type]() []T
	GetI() interface{}
	Loc(p []int) Vector
	Err() error
	Copy() Vector
}

// NewVec ...
func NewVec(data interface{}) Vector {
	switch t := data.(type) {
	case []int:
		return IntVector{
			obs:  t,
			size: len(t),
		}

	case []string:
		return StrVector{
			obs:  t,
			size: len(t),
		}
	default:
		return StrVector{
			err: fmt.Errorf("wrong data type, expected []int or []string, got %T", t),
		}
	}
}

// IntVector ... ---------------------------------------------------------------
type IntVector struct {
	obs   []int
	index []int
	size  int
	err   error
}

// Size ...
func (v IntVector) Size() int {
	return v.size
}

// Get ...
func (v IntVector) GetI() interface{} {
	new := make([]int, v.Size())
	copy(new, v.obs)
	return new
}

func (v IntVector) Get() []int {
	return v.obs
}

// Err  ...
func (v IntVector) Err() error {
	return v.err
}

// Loc ...
func (v IntVector) Loc(p []int) Vector {
	new := make([]int, len(p))
	for ix, val := range p {
		if val >= v.Size() {
			return IntVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
		}
		new[ix] = v.obs[val]
	}
	return IntVector{
		obs:  new,
		size: len(p),
	}
}

// Copy ...
func (v IntVector) Copy() Vector {

	new := make([]int, v.size)
	cp := copy(new, v.obs)
	if cp != v.size {
		return StrVector{
			err: fmt.Errorf("copy returned a wrong number of elements, expected: %v, got:%v", v.size, cp),
		}
	}

	return IntVector{
		obs:   new,
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
		index: v.index,
		size:  v.size,
	}
}

// Find ...
func (v IntVector) Find(f func(v int) bool) []bool {
	new := make([]bool, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
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
	index   []int
	size    int
	inverse map[string][]int //?inverse index
	err     error
}

// Size ...
func (v StrVector) Size() int {
	return v.size
}

// Get ...
func (v StrVector) GetI() interface{} {
	return v.obs
}

func (v StrVector) Get() []string {
	new := make([]string, v.Size())
	copy(new, v.obs)
	return new
}

// Err  ...
func (v StrVector) Err() error {
	return v.err
}

// Loc ...
func (v StrVector) Loc(p []int) Vector {
	new := make([]string, len(p))
	for ix, val := range p {
		if val >= v.Size() {
			return StrVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
		}
		new[ix] = v.obs[val]
	}
	return StrVector{
		obs:  new,
		size: len(new),
	}
}

// Copy ...
func (v StrVector) Copy() Vector {

	new := make([]string, v.size)
	cp := copy(new, v.obs)
	if cp != v.size {
		return StrVector{
			err: fmt.Errorf("copy returned a wrong number of elements, expected: %v, got:%v", v.size, cp),
		}
	}

	return StrVector{
		obs:   new,
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
		index: v.index,
		size:  v.size,
	}
}

// Find ...
func (v StrVector) Find(f func(v string) bool) []bool {
	new := make([]bool, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
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
