package vec

import (
	"fmt"
	"strconv"
)

// IntVector ... ---------------------------------------------------------------
type IntVector struct {
	obs  []int
	na   Set
	hash struct {
		lookup map[int]int
		size   int
	}
	index map[int][]int
	size  int
	err   error
}

// Size ...
func (v IntVector) Size() int {
	return v.size
}

// StrVector ... ---------------------------------------------------------------
type StrVector struct {
	obs  []string
	na   Set
	hash struct {
		lookup map[string]int
		size   int
	}
	index   map[string][]int
	size    int
	inverse map[string][]int //! inverse index
	err     error
}

// Size ...
func (v StrVector) Size() int {
	return v.size
}

// ToStr ...
func (v IntVector) ToStr() StrVector {

	if v.Err() != nil {
		return StrVector{
			err: fmt.Errorf("ToStr - Vector already contains an error: %w", v.Err()),
		}
	}

	data := make([]string, v.size)

	for ix, val := range v.obs {
		data[ix] = strconv.Itoa(val)
	}

	return StrVector{
		obs:  data,
		na:   v.na,
		size: v.size,
	}
}

// ToStr ...
func (v StrVector) ToStr() StrVector {

	return v

}
