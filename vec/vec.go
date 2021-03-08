// Package vec implements the Vector type, an underlying structure on which
// each Table is build.
package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// Vector ...
type Vector interface {
	Size() int
	Int() IntVector
	Str() StrVector
	//Get[T vector_type]() []T
	GetI() (interface{}, Set)
	Loc(p []int) Vector
	Check(interface{}) ([]bool, error)
	Filter(interface{}) Vector
	Mask([]bool) Vector
	Err() error
	setError(error) Vector
	Hash() Vector
	IsHashed() bool
	GetHashVals() ([]int, int)
	SetHash(Vector) Vector
	Copy() Vector
	Sort() Vector
	Order() []int
	Mutate(interface{}) Vector
	Group() Vector
	//TODO: remove after testing
	Elem(int) (interface{}, bool)
	ToStr() StrVector
}

// NewErrVec ...
func NewErrVec(err error) Vector {
	return StrVector{
		err: err,
	}
}

// NewVec ...
func NewVec(data interface{}, nas ...Set) Vector {

	var na Set
	if len(nas) == 0 {
		na = make(Set)
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
			err: fmt.Errorf("%w wrong data type, expected `[]int` / `[]string`, got `%T`", utils.ErrParamType, t),
		}
	}
}
