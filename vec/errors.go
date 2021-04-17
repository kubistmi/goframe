package vec

import (
	"fmt"
)

// NewErrVector prepares a new Vector of the selected type with just the error
func NewErrVec(err error, kind Datatype) Vector {
	switch kind {
	case IntType:
		return IntVector{
			data: []int{},
			na:   NewSet(-1),
			err:  err,
		}
	case StrType:
		return StrVector{
			data: []string{},
			na:   NewSet(-1),
			err:  err,
		}
	case BoolType:
		return BoolVector{
			data: []bool{},
			na:   NewSet(-1),
			err:  err,
		}
	}
	return NewErrVec(fmt.Errorf("not implemented"), StrType)
}

// IntVector implementations ---------------------------------------------------

// Err returns the error from previous operations
func (v IntVector) Err() error {
	return v.err
}

// newErrIntVector prepares a new IntVector with just the error
func newErrIntVec(err error) IntVector {
	return IntVector{
		data: []int{},
		na:   NewSet(-1),
		err:  err,
	}
}

// StrVector implementations ---------------------------------------------------

// Err returns the error from previous operations
func (v StrVector) Err() error {
	return v.err
}

// newErrVector prepares a new StrVector with just the error
func newErrStrVec(err error) StrVector {
	return StrVector{
		data: []string{},
		na:   NewSet(-1),
		err:  err,
	}
}

// StrVector implementations ---------------------------------------------------

// Err returns the error from previous operations
func (v BoolVector) Err() error {
	return v.err
}

// newErrVector prepares a new StrVector with just the error
func newErrBoolVec(err error) BoolVector {
	return BoolVector{
		data: []bool{},
		na:   NewSet(-1),
		err:  err,
	}
}
