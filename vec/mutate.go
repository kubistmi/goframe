package vec

import (
	"fmt"
)

// Mutate ...
//TODO: handle NAs
func (v IntVector) Mutate(f interface{}) Vector {
	fun, ok := f.(func(v int) int)
	if !ok {
		return StrVector{err: fmt.Errorf("wrong function, expected: `func(int) int`, got `%T`", f)}
	}

	new := make([]int, v.Size())
	for ix, val := range v.obs {
		new[ix] = fun(val)
	}
	return IntVector{
		obs:  new,
		na:   v.na,
		hash: v.hash,
		size: v.size,
	}
}

// Mutate ...
//TODO: handle NAs
func (v StrVector) Mutate(f interface{}) Vector {
	fun, ok := f.(func(v string) string)
	if !ok {
		return StrVector{err: fmt.Errorf("wrong function, expected: `func(string) string`, got `%T`", f)}
	}

	new := make([]string, v.Size())
	for ix, val := range v.obs {
		new[ix] = fun(val)
	}
	return StrVector{
		obs:  new,
		na:   v.na,
		hash: v.hash,
		size: v.size,
	}
}
