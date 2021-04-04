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
			return NewErrVec(fmt.Errorf("interface `f` contains error: %w", err), v.Type())
		}
		return NewErrVec(fmt.Errorf("%w parameter f, expected: `func(int) int`, got `%T`", utils.ErrParamType, f), v.Type())
	}

	new := make([]int, v.Size())
	newNA := make(Set, v.na.Size())
	var na bool
	for ix, val := range v.data {
		new[ix], na = fun(val, v.na.Get(ix))
		if na {
			newNA.Set(ix)
		}
	}
	return IntVector{
		data: new,
		na:   newNA,
		hash: v.hash,
	}
}

// Mutate ...
func (v StrVector) Mutate(f interface{}) Vector {
	fun, ok := f.(func(string, bool) (string, bool))
	if !ok {
		if err, ok := f.(error); ok {
			return NewErrVec(fmt.Errorf("interface `f` contains error: %w", err), v.Type())
		}
		return NewErrVec(fmt.Errorf("%w parameter f, expected: `func(string) string`, got `%T`", utils.ErrParamType, f), v.Type())
	}

	new := make([]string, v.Size())
	newNA := make(Set, v.na.Size())
	var na bool
	for ix, val := range v.data {
		new[ix], na = fun(val, v.na.Get(ix))
		if na {
			newNA.Set(ix)
		}
	}
	return StrVector{
		data: new,
		na:   newNA,
		hash: v.hash,
	}
}
