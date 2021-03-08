package vec

import (
	"fmt"
	"strings"

	"github.com/kubistmi/goframe/utils"
)

// Str ...
func (v IntVector) Str() StrVector {
	return v.setStrError(fmt.Errorf("%w cant type switch, expected: `StrVector`, got: `%T`", utils.ErrParamType, v))
}

// Str ...
func (v StrVector) Str() StrVector {
	return v
}

// Int ...
func (v IntVector) Int() IntVector {
	return v
}

// Int ...
func (v StrVector) Int() IntVector {
	return v.setIntError(fmt.Errorf("%w cant type switch, expected: `IntVector`, got: `%T`", utils.ErrParamType, v))
}

// JoinStr ...
func (v StrVector) JoinStr(s interface{}) StrVector {
	new := make([]string, v.Size())
	na := v.na.Copy()

	switch s := s.(type) {
	case string:
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				new[ix] = val + s
			}
		}
	case []string:
		if len(s) != v.Size() {
			return v.setStrError(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				new[ix] = val + s[ix]
			}
		}
	case StrVector:
		if s.Size() != v.Size() {
			return v.setStrError(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = ""
			} else {
				new[ix] = val + s.obs[ix]
			}
		}
	}

	return StrVector{
		obs:  new,
		na:   na,
		size: v.Size(),
	}
}

// Rep ...
func (v StrVector) Rep(s interface{}) StrVector {
	new := make([]string, v.Size())
	na := v.na.Copy()

	switch s := s.(type) {
	case int:
		if s <= 0 {
			return v.setStrError(fmt.Errorf("%w non-positive number of repeats `s`, expected `s > 0`, got %v", utils.ErrParamVal, s))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				new[ix] = strings.Repeat(val, s)
			}
		}
	case []int:
		if len(s) != v.Size() {
			return v.setStrError(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[ix] <= 0 {
					return v.setStrError(fmt.Errorf("%w non-positive number of repeats `s` at position %v, expected `s > 0`, got %v", utils.ErrParamVal, ix, s[ix]))
				}
				new[ix] = strings.Repeat(val, s[ix])
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return v.setStrError(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = ""
			} else {
				if s.obs[ix] < 0 {
					return v.setStrError(fmt.Errorf("%w non-positive number of repeats `s` at position %v, expected `s > 0`, got %v", utils.ErrParamVal, ix, s.obs[ix]))
				}
				new[ix] = strings.Repeat(val, s.obs[ix])
			}
		}
	}

	return StrVector{
		obs:  new,
		na:   na,
		size: v.Size(),
	}
}

// Sub ...
func (v StrVector) Sub(s interface{}) StrVector {
	new := make([]string, v.Size())
	na := v.na.Copy()

	switch s := s.(type) {
	case [2]int:
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[0] < 0 || s[1] <= s[0] || s[1] > len(val) {
					return v.setStrError(fmt.Errorf("%w wrong range specified in `s`, string size: `%v`, got `[%v, %v]`", utils.ErrParamVal, len(val), s[0], s[1]))
				}
				new[ix] = val[s[0]:s[1]]
			}
		}
	case [][2]int:
		if len(s) != v.Size() {
			return v.setStrError(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[ix][0] < 0 || s[ix][1] <= s[ix][0] || s[ix][1] > len(val) {
					return v.setStrError(fmt.Errorf("%w wrong range specified in `s`, string size: `%v`, got `[%v, %v]`", utils.ErrParamVal, len(val), s[ix][0], s[ix][1]))
				}
				new[ix] = val[s[ix][0]:s[ix][1]]
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return v.setStrError(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = ""
			} else {
				new[ix] = strings.Repeat(val, s.obs[ix])
			}
		}
	}

	return StrVector{
		obs:  new,
		na:   na,
		size: v.Size(),
	}
}

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
	na := v.na.Copy()

	switch s := s.(type) {
	case int:
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = 0
			} else {
				new[ix] = doIntOp(val, s, kind)
			}
		}
	case []int:
		if len(s) != v.Size() {
			return v.setIntError(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = 0
			} else {
				new[ix] = doIntOp(val, s[ix], kind)
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return v.setIntError(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = 0
			} else {
				new[ix] = doIntOp(val, s.obs[ix], kind)
			}
		}
	}

	return IntVector{
		obs:  new,
		na:   na,
		size: v.Size(),
	}

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
