package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// Mutate ...
func (v IntVector) Mutate(f interface{}) Vector {
	fun, ok := f.(func(int, bool) (int, bool))
	if !ok {
		if err, ok := f.(error); ok {
			return v.setError(fmt.Errorf("interface `f` contains error: %w", err))
		}
		return v.setError(fmt.Errorf("%w parameter f, expected: `func(int) int`, got `%T`", utils.ErrParamType, f))
	}

	new := make([]int, v.Size())
	newNA := make(Set, len(v.na))
	var na bool
	for ix, val := range v.obs {
		new[ix], na = fun(val, v.na.Get(ix))
		if na {
			newNA = newNA.Set(ix)
		}
	}
	return IntVector{
		obs:  new,
		na:   newNA,
		hash: v.hash,
		size: v.size,
	}
}

// Mutate ...
func (v StrVector) Mutate(f interface{}) Vector {
	fun, ok := f.(func(string, bool) (string, bool))
	if !ok {
		if err, ok := f.(error); ok {
			return v.setError(fmt.Errorf("interface `f` contains error: %w", err))
		}
		return v.setError(fmt.Errorf("%w parameter f, expected: `func(string) string`, got `%T`", utils.ErrParamType, f))
	}

	new := make([]string, v.Size())
	newNA := make(Set, len(v.na))
	var na bool
	for ix, val := range v.obs {
		new[ix], na = fun(val, v.na.Get(ix))
		if na {
			newNA = newNA.Set(ix)
		}
	}
	return StrVector{
		obs:  new,
		na:   newNA,
		hash: v.hash,
		size: v.size,
	}
}
