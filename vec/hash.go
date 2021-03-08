package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// Hash ...
func (v IntVector) Hash() Vector {

	nhash := make(map[int]int)

	i := 1
	for _, val := range v.obs {
		if _, ok := nhash[val]; !ok {
			nhash[val] = i
			i++
		}
	}

	nhix := struct {
		lookup map[int]int
		size   int
	}{nhash, i}

	v.hash = nhix
	return v
}

// Hash ...
func (v StrVector) Hash() Vector {
	nhash := make(map[string]int)

	i := 1
	for _, val := range v.obs {
		if _, ok := nhash[val]; !ok {
			nhash[val] = i
			i++
		}
	}

	nhix := struct {
		lookup map[string]int
		size   int
	}{nhash, i}

	v.hash = nhix
	return v
}

// GetHash ...
func (v IntVector) GetHash(l int) int {
	return v.hash.lookup[l]
}

// GetHash ...
func (v StrVector) GetHash(l string) int {
	return v.hash.lookup[l]
}

// IsHashed ...
func (v IntVector) IsHashed() bool {
	return len(v.hash.lookup) > 0
}

// IsHashed ...
func (v StrVector) IsHashed() bool {
	return len(v.hash.lookup) > 0
}

// GetHashVals ...
func (v IntVector) GetHashVals() ([]int, int) {
	out := make([]int, v.size)
	for ix, val := range v.obs {
		out[ix] = v.hash.lookup[val]
	}
	return out, v.hash.size
}

// GetHashVals ...
func (v StrVector) GetHashVals() ([]int, int) {
	out := make([]int, v.size)
	for ix, val := range v.obs {
		out[ix] = v.hash.lookup[val]
	}
	return out, v.hash.size
}

// SetHash ...
func (v StrVector) SetHash(ri Vector) Vector {

	r, ok := ri.(StrVector)
	if !ok {
		return v.setError(fmt.Errorf("%w parameter ri, expected: `%T`, got: `%T`", utils.ErrParamType, v, ri))
	}

	var new struct {
		lookup map[string]int
		size   int
	}

	new.lookup = make(map[string]int, len(r.hash.lookup))
	for ix, val := range r.hash.lookup {
		new.lookup[ix] = val
	}
	new.size = r.hash.size

	v.hash = new
	return v
}

// SetHash ...
func (v IntVector) SetHash(ri Vector) Vector {
	r, ok := ri.(IntVector)
	if !ok {
		return v.setError(fmt.Errorf("%w parameter ri, expected: `%T`, got: `%T`", utils.ErrParamType, v, ri))
	}

	var new struct {
		lookup map[int]int
		size   int
	}

	new.lookup = make(map[int]int, len(r.hash.lookup))
	for ix, val := range r.hash.lookup {
		new.lookup[ix] = val
	}
	new.size = r.hash.size

	v.hash = new
	return v
}
