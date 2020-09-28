package vec

import "fmt"

// Vector ...
type Vector interface {
	Size() int
	//Get[T vector_type]() []T
	GetI() (interface{}, map[int]bool)
	Loc(p []int) Vector
	Err() error
	Copy() Vector
	Hash() Vector
	GetHashVals() ([]int, int)
	IsHashed() bool
	//TODO: remove after testing
	Elem(int) (interface{}, bool)
}

// NewErrVec ...
func NewErrVec(err error) Vector {
	return StrVector{
		err: err,
	}
}

// NewVec ...
func NewVec(data interface{}, nas ...map[int]bool) Vector {

	var na map[int]bool
	if len(nas) == 0 {
		na = make(map[int]bool)
	} else {
		na = nas[0]
	}

	switch t := data.(type) {
	case []int:
		new := make([]int, len(t))
		copy(new, t)
		return IntVector{
			obs:  new,
			na:   na,
			size: len(t),
		}

	case []string:
		new := make([]string, len(t))
		copy(new, t)
		return StrVector{
			obs:  new,
			na:   na,
			size: len(t),
		}
	default:
		return StrVector{
			err: fmt.Errorf("wrong data type, expected []int or []string, got %T", t),
		}
	}
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
