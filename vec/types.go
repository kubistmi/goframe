package vec

import (
	"fmt"
	"strconv"

	"github.com/kubistmi/goframe/utils"
)

// Datatype describes the type of the Vector
type Datatype int

const (
	//StringType describes a Vector type based on []string
	StrType Datatype = iota

	//IntType describes a Vector type based on []int
	IntType

	//NumType describes a Vector type based on []float
	NumType

	//BoolType describes a Vector type based on []bool
	BoolType

	//DateType describes a Vector type based on []time.Time
	DateType
)

// IntVector implementations ---------------------------------------------------

// IntVector is an data structure build on top of []int
type IntVector struct {
	data []int
	na   NA
	hash struct {
		lookup map[int]int
		size   int
	}
	index map[int][]int
	err   error
}

// NewIntVector creates an IntVector and populates it with a copy of used values
func NewIntVec(d []int, n NA) IntVector {
	new := make([]int, len(d))
	copy(new, d)
	newna := n.CopyNA()

	return IntVector{
		data: new,
		na:   newna,
	}
}

// unsageIntVector creates a new IntVector and pushes the parameters without copying
func unsafeIntVec(d []int, na NA) IntVector {
	return IntVector{
		data: d,
		na:   na,
		err:  nil,
	}
}

// Size returns the length of the underlying slice ~ len([]T)
func (v IntVector) Size() int {
	return len(v.data)
}

// Type returns the Datatype of the Vector
func (v IntVector) Type() Datatype {
	return IntType
}

// ToStr changes the type of the Vector to StrVector
func (v IntVector) ToStr() StrVector {

	if v.Err() != nil {
		return newErrStrVec(fmt.Errorf("ToStr - %w %w", utils.ErrAlreadyErr, v.Err()))
	}

	data := make([]string, v.Size())

	for ix, val := range v.data {
		if v.na.Get(ix) {
			data[ix] = ""
		} else {
			data[ix] = strconv.Itoa(val)
		}
	}

	return NewStrVec(data, v.na.CopyNA())
}

// Str attempts to type switch the Vector to StrVector
func (v IntVector) Str() StrVector {
	return newErrStrVec(fmt.Errorf("%w cant type switch, expected: `StrVector`, got: `%T`", utils.ErrParamType, v))
}

// Int attempts to type switch the Vector to IntVector
func (v IntVector) Int() IntVector {
	return v
}

// StrVector implementations ---------------------------------------------------

// StrVector is an data structure build on top of []string
type StrVector struct {
	data []string
	na   NA
	hash struct {
		lookup map[string]int
		size   int
	}
	index   map[string][]int
	inverse map[string][]int //! inverse index
	err     error
}

// NewStrVector creates an StrVector and populates it with a copy of used values
func NewStrVec(d []string, n NA) StrVector {
	new := make([]string, len(d))
	copy(new, d)
	newna := n.CopyNA()

	return StrVector{
		data: new,
		na:   newna,
	}
}

// unsageStrVector creates a new StrVector and pushes the parameters without copying
func unsafeStrVec(d []string, na NA) StrVector {
	return StrVector{
		data: d,
		na:   na,
		err:  nil,
	}
}

// Size returns the length of the underlying slice ~ len([]T)
func (v StrVector) Size() int {
	return len(v.data)
}

// Type returns the Datatype of the Vector
func (v StrVector) Type() Datatype {
	return StrType
}

// ToStr changes the type of the Vector to StrVector
func (v StrVector) ToStr() StrVector {
	return v
}

// Str attempts to type switch the Vector to StrVector
func (v StrVector) Str() StrVector {
	return v
}

// Int attempts to type switch the Vector to IntVector
func (v StrVector) Int() IntVector {
	return newErrIntVec(fmt.Errorf("%w cant type switch, expected: `IntVector`, got: `%T`", utils.ErrParamType, v))
}
