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
	Type() Datatype
	Err() error
	GetI() (interface{}, NA, error)
	Copy() Vector
	Int() IntVector
	Str() StrVector
	ToStr() StrVector
	Loc(p []int) Vector
	Check(interface{}) ([]bool, error)
	Filter(interface{}) Vector
	Mask([]bool) Vector
	Hash() Vector
	IsHashed() bool
	GetHashVals() ([]int, int)
	SetHash(Vector) Vector
	Sort() Vector
	Order() []int
	Mutate(interface{}) Vector
	Group() Vector
	Bool() BoolVector

	//! concepts
	ElemI(i int) (interface{}, bool)
	// Is(val interface{}) func() (bool, error)
}

// NewVec ...
func NewVec(data interface{}, na NA) Vector {

	if na == nil {
		na = NewNA(-1)
	}

	switch t := data.(type) {
	case []int:
		new := make([]int, len(t))
		copy(new, t)
		return IntVector{
			data: new,
			na:   na,
		}

	case []string:
		new := make([]string, len(t))
		copy(new, t)
		return StrVector{
			data: new,
			na:   na,
		}
	case []bool:
		new := make([]bool, len(t))
		copy(new, t)
		return BoolVector{
			data: new,
			na:   na,
		}
	default:
		return StrVector{
			err: fmt.Errorf("%w wrong data type, expected `[]int` / `[]string`, got `%T`", utils.ErrParamType, t),
		}
	}
}
