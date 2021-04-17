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

// NewIntVec creates an IntVector and populates it with a copy of used values
func NewIntVec(d []int, n NA) IntVector {
	new := make([]int, len(d))
	copy(new, d)
	newna := n.CopyNA()

	return IntVector{
		data: new,
		na:   newna,
	}
}

// unsafeIntVector creates a new IntVector and pushes the parameters without copying
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

// Bool attempts to type switch the Vector to BoolVector
func (v IntVector) Bool() BoolVector {
	return newErrBoolVec(fmt.Errorf("%w cant type switch, expected: `BoolVector`, got: `%T`", utils.ErrParamType, v))
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

// NewStrVec creates an StrVector and populates it with a copy of used values
func NewStrVec(d []string, n NA) StrVector {
	new := make([]string, len(d))
	copy(new, d)
	newna := n.CopyNA()

	return StrVector{
		data: new,
		na:   newna,
	}
}

// unsafeStrVector creates a new StrVector and pushes the parameters without copying
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

func (v StrVector) Bool() BoolVector {
	return newErrBoolVec(fmt.Errorf("%w cant type switch, expected: `BoolVector`, got: `%T`", utils.ErrParamType, v))
}

// BoolVector implementations ---------------------------------------------------

// BoolVector is an data structure build on top of []bool
type BoolVector struct {
	data []bool
	na   NA
	hash struct {
		lookup map[bool]int
		size   int
	}
	index map[bool][]int
	err   error
}

// NewBoolVec creates an BoolVector and populates it with a copy of used values
func NewBoolVec(d []bool, n NA) BoolVector {
	new := make([]bool, len(d))
	copy(new, d)
	newna := n.CopyNA()

	return BoolVector{
		data: new,
		na:   newna,
	}
}

// unsafeBoolVector creates a new BoolVector and pushes the parameters without copying
func unsafeBoolVec(d []bool, na NA) BoolVector {
	return BoolVector{
		data: d,
		na:   na,
		err:  nil,
	}
}

// Size returns the length of the underlying slice ~ len([]T)
func (v BoolVector) Size() int {
	return len(v.data)
}

// Type returns the Datatype of the Vector
func (v BoolVector) Type() Datatype {
	return BoolType
}

// ToStr changes the type of the Vector to StrVector
func (v BoolVector) ToStr() StrVector {

	if v.Err() != nil {
		return newErrStrVec(fmt.Errorf("ToStr - %w %w", utils.ErrAlreadyErr, v.Err()))
	}

	data := make([]string, v.Size())

	for ix, val := range v.data {
		if v.na.Get(ix) {
			data[ix] = ""
		} else {
			data[ix] = strconv.FormatBool(val)
		}
	}

	return NewStrVec(data, v.na.CopyNA())
}

// Str attempts to type switch the Vector to StrVector
func (v BoolVector) Str() StrVector {
	return newErrStrVec(fmt.Errorf("%w cant type switch, expected: `StrVector`, got: `%T`", utils.ErrParamType, v))
}

// Int attempts to type switch the Vector to IntVector
func (v BoolVector) Int() IntVector {
	return newErrIntVec(fmt.Errorf("%w cant type switch, expected: `IntVector`, got: `%T`", utils.ErrParamType, v))
}

func (v BoolVector) Bool() BoolVector {
	return v
}
