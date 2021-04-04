package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// Add ...
func (v IntVector) Add(s interface{}) IntVector {
	return runIntOp(v, s, "add")
}

// Mult ...
func (v IntVector) Mult(s interface{}) IntVector {
	return runIntOp(v, s, "mult")
}

// Sub ...
func (v IntVector) Sub(s interface{}) IntVector {
	return runIntOp(v, s, "sub")
}

// Div ...
func (v IntVector) Div(s interface{}) IntVector {
	return runIntOp(v, s, "div")
}

func runIntOp(v IntVector, s interface{}, kind string) IntVector {
	new := make([]int, v.Size())
	na := v.na.CopyNA()

	switch s := s.(type) {
	case int:
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = 0
				na.Set(ix)
			} else {
				new[ix] = doIntOp(val, s, kind)
			}
		}
	case []int:
		if len(s) != v.Size() {
			return newErrIntVec(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = 0
				na.Set(ix)
			} else {
				new[ix] = doIntOp(val, s[ix], kind)
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return newErrIntVec(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) || s.na.Get(ix) {
				new[ix] = 0
				na.Set(ix)
			} else {
				new[ix] = doIntOp(val, s.data[ix], kind)
			}
		}
	}

	return NewIntVec(new, na)
}

func doIntOp(x, y int, kind string) int {
	if kind == "add" {
		return x + y
	} else if kind == "sub" {
		return x - y
	} else if kind == "mult" {
		return x * y
	} else if kind == "div" {
		return x / y
	}
	return x
}
