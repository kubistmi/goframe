package vec

import (
	"fmt"
	"strconv"

	"github.com/kubistmi/goframe/utils"
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
func (v IntVector) Size() int {
	return v.size
}

// Size ...
func (v StrVector) Size() int {
	return v.size
}

// ToStr ...
func (v IntVector) ToStr() StrVector {

	if v.Err() != nil {
		return StrVector{
			err: fmt.Errorf("ToStr - %w %w", utils.ErrAlreadyErr, v.Err()),
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
