package vec

import (
	"fmt"
	"strings"

	"github.com/kubistmi/goframe/utils"
)

// JoinStr ...
func (v StrVector) JoinStr(s interface{}) StrVector {
	new := make([]string, v.Size())
	na := v.na.CopyNA()

	switch s := s.(type) {
	case string:
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				new[ix] = val + s
			}
		}
	case []string:
		if len(s) != v.Size() {
			return newErrStrVec(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				new[ix] = val + s[ix]
			}
		}
	case StrVector:
		if s.Size() != v.Size() {
			return newErrStrVec(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = ""
			} else {
				new[ix] = val + s.data[ix]
			}
		}
	}

	return StrVector{
		data: new,
		na:   na,
	}
}

// Rep ...
func (v StrVector) Rep(s interface{}) StrVector {
	new := make([]string, v.Size())
	na := v.na.CopyNA()

	switch s := s.(type) {
	case int:
		if s <= 0 {
			return newErrStrVec(fmt.Errorf("%w non-positive number of repeats `s`, expected `s > 0`, got %v", utils.ErrParamVal, s))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				new[ix] = strings.Repeat(val, s)
			}
		}
	case []int:
		if len(s) != v.Size() {
			return newErrStrVec(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[ix] <= 0 {
					return newErrStrVec(fmt.Errorf("%w non-positive number of repeats `s` at position %v, expected `s > 0`, got %v", utils.ErrParamVal, ix, s[ix]))
				}
				new[ix] = strings.Repeat(val, s[ix])
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return newErrStrVec(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = ""
			} else {
				if s.data[ix] < 0 {
					return newErrStrVec(fmt.Errorf("%w non-positive number of repeats `s` at position %v, expected `s > 0`, got %v", utils.ErrParamVal, ix, s.data[ix]))
				}
				new[ix] = strings.Repeat(val, s.data[ix])
			}
		}
	}

	return StrVector{
		data: new,
		na:   na,
	}
}

// Sub ...
func (v StrVector) Sub(s interface{}) StrVector {
	new := make([]string, v.Size())
	na := v.na.CopyNA()

	switch s := s.(type) {
	case [2]int:
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[0] < 0 || s[1] <= s[0] || s[1] > len(val) {
					return newErrStrVec(fmt.Errorf("%w wrong range specified in `s`, string size: `%v`, got `[%v, %v]`", utils.ErrParamVal, len(val), s[0], s[1]))
				}
				new[ix] = val[s[0]:s[1]]
			}
		}
	case [][2]int:
		if len(s) != v.Size() {
			return newErrStrVec(fmt.Errorf("%w wrong size of slice `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), len(s)))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[ix][0] < 0 || s[ix][1] <= s[ix][0] || s[ix][1] > len(val) {
					return newErrStrVec(fmt.Errorf("%w wrong range specified in `s`, string size: `%v`, got `[%v, %v]`", utils.ErrParamVal, len(val), s[ix][0], s[ix][1]))
				}
				new[ix] = val[s[ix][0]:s[ix][1]]
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return newErrStrVec(fmt.Errorf("%w wrong size of IntVector `s`, expected: `%v`, got `%v`", utils.ErrParamVal, v.Size(), s.Size()))
		}
		for ix, val := range v.data {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = ""
			} else {
				new[ix] = strings.Repeat(val, s.data[ix])
			}
		}
	}

	return StrVector{
		data: new,
		na:   na,
	}
}
