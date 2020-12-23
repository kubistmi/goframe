// Package vec implements the Vector type, an underlying structure on which
// each Table is build.
package vec

import "fmt"

// Vector ...
type Vector interface {
	Size() int
	//Get[T vector_type]() []T
	GetI() (interface{}, Set)
	Loc(p []int) Vector
	Check(interface{}) ([]bool, error)
	Filter(interface{}) Vector
	Mask([]bool) Vector
	AssignM([]int, interface{}, Set) Vector
	AssignI(int, interface{}, bool) Vector
	Err() error
	Hash() Vector
	IsHashed() bool
	GetHashVals() ([]int, int)
	Copy() Vector
	Sort() Vector
	Order() []int
	Mutate(interface{}) Vector
	Group() Vector
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
			err: fmt.Errorf("wrong data type, expected []int or []string, got %T", t),
		}
	}
}
