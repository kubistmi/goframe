package vec

import (
	"fmt"
	"strings"
)

// Str ...
func (v IntVector) Str() StrVector {
	return StrVector{err: fmt.Errorf("cannot type switch, expected: `StrVector`, got: %T", v)}
}

// Str ...
func (v StrVector) Str() StrVector {
	return v
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
			return StrVector{err: fmt.Errorf("wrong slice size, expected: %v, got %v", v.Size(), len(s))}
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
			return StrVector{err: fmt.Errorf("wrong IntVector size, expected: %v, got %v", v.Size(), s.Size())}
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
		if s < 0 {
			return StrVector{err: fmt.Errorf("negative number of repeats, got %v", s)}
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
			return StrVector{err: fmt.Errorf("wrong slice size, expected: %v, got %v", v.Size(), len(s))}
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[ix] < 0 {
					return StrVector{err: fmt.Errorf("negative number of repeats at position %v", ix)}
				}
				new[ix] = strings.Repeat(val, s[ix])
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return StrVector{err: fmt.Errorf("wrong IntVector size, expected: %v, got %v", v.Size(), s.Size())}
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) || s.na.Get(ix) {
				na.Set(ix)
				new[ix] = ""
			} else {
				if s.obs[ix] < 0 {
					return StrVector{err: fmt.Errorf("negative number of repeats at position %v", ix)}
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
					return StrVector{err: fmt.Errorf("wrong sub specification, string size: %v, got [%v, %v]", len(val), s[0], s[1])}
				}
				new[ix] = val[s[0]:s[1]]
			}
		}
	case [][2]int:
		if len(s) != v.Size() {
			return StrVector{err: fmt.Errorf("wrong slice size, expected: %v, got %v", v.Size(), len(s))}
		}
		for ix, val := range v.obs {
			if v.na.Get(ix) {
				new[ix] = ""
			} else {
				if s[ix][0] < 0 || s[ix][1] <= s[ix][0] || s[ix][1] > len(val) {
					return StrVector{err: fmt.Errorf("wrong sub specification, string size: %v, got [%v, %v]", len(val), s[ix][0], s[ix][1])}
				}
				new[ix] = val[s[ix][0]:s[ix][1]]
			}
		}
	case IntVector:
		if s.Size() != v.Size() {
			return StrVector{err: fmt.Errorf("wrong IntVector size, expected: %v, got %v", v.Size(), s.Size())}
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

// Int ...
func (v IntVector) Int() IntVector {
	return v
}

// Int ...
func (v StrVector) Int() IntVector {
	return IntVector{err: fmt.Errorf("cannot type switch, expected: `IntVector`, got: %T", v)}
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
			return IntVector{err: fmt.Errorf("wrong slice size, expected: %v, got %v", v.Size(), len(s))}
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
			return IntVector{err: fmt.Errorf("wrong IntVector size, expected: %v, got %v", v.Size(), s.Size())}
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
