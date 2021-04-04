package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// IntVector implementations ---------------------------------------------------

type intHash struct {
	lookup map[int]int
	size   int
}

func (h intHash) Copy() intHash {
	new := make(map[int]int, len(h.lookup))
	for ix, val := range h.lookup {
		new[ix] = val
	}
	return intHash{
		lookup: new,
		size:   h.size,
	}
}

// Hash ...
func (v IntVector) Hash() Vector {

	nhash := make(map[int]int)

	i := 1
	for _, val := range v.data {
		if _, ok := nhash[val]; !ok {
			nhash[val] = i
			i++
		}
	}

	v.hash = intHash{nhash, i}
	return v
}

// GetHash ...
func (v IntVector) GetHash(l int) int {
	return v.hash.lookup[l]
}

// IsHashed ...
func (v IntVector) IsHashed() bool {
	return len(v.hash.lookup) > 0
}

// GetHashVals ...
func (v IntVector) GetHashVals() ([]int, int) {
	out := make([]int, v.Size())
	for ix, val := range v.data {
		out[ix] = v.hash.lookup[val]
	}
	return out, v.hash.size
}

// SetHash ...
func (v IntVector) SetHash(ri Vector) Vector {
	r, ok := ri.(IntVector)
	if !ok {
		return NewErrVec(fmt.Errorf("%w parameter ri, expected: `%T`, got: `%T`", utils.ErrParamType, v, ri), v.Type())
	}

	var new intHash

	new.lookup = make(map[int]int, len(r.hash.lookup))
	for ix, val := range r.hash.lookup {
		new.lookup[ix] = val
	}
	new.size = r.hash.size

	v.hash = new
	return v
}

// StrVector implementations ---------------------------------------------------

type strHash struct {
	lookup map[string]int
	size   int
}

func (h strHash) Copy() strHash {
	new := make(map[string]int, len(h.lookup))
	for ix, val := range h.lookup {
		new[ix] = val
	}
	return strHash{
		lookup: new,
		size:   h.size,
	}
}

// Hash ...
func (v StrVector) Hash() Vector {
	nhash := make(map[string]int)

	i := 1
	for _, val := range v.data {
		if _, ok := nhash[val]; !ok {
			nhash[val] = i
			i++
		}
	}

	nhix := strHash{nhash, i}

	v.hash = nhix
	return v
}

// GetHash ...
func (v StrVector) GetHash(l string) int {
	return v.hash.lookup[l]
}

// IsHashed ...
func (v StrVector) IsHashed() bool {
	return len(v.hash.lookup) > 0
}

// GetHashVals ...
func (v StrVector) GetHashVals() ([]int, int) {
	out := make([]int, v.Size())
	for ix, val := range v.data {
		out[ix] = v.hash.lookup[val]
	}
	return out, v.hash.size
}

// SetHash ...
func (v StrVector) SetHash(ri Vector) Vector {

	r, ok := ri.(StrVector)
	if !ok {
		return NewErrVec(fmt.Errorf("%w parameter ri, expected: `%T`, got: `%T`", utils.ErrParamType, v, ri), v.Type())
	}

	var new strHash

	new.lookup = make(map[string]int, len(r.hash.lookup))
	for ix, val := range r.hash.lookup {
		new.lookup[ix] = val
	}
	new.size = r.hash.size

	v.hash = new
	return v
}
