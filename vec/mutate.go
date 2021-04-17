package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

// IntVector implementations ---------------------------------------------------

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
	newna := make(Set, v.na.Size())
	var na bool
	for ix, val := range v.data {
		new[ix], na = fun(val, v.na.Get(ix))
		newna.Setif(na, ix)
	}
	return NewIntVec(new, newna)
}

// StrVector implementations ---------------------------------------------------

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
	newna := make(Set, v.na.Size())
	var na bool
	for ix, val := range v.data {
		new[ix], na = fun(val, v.na.Get(ix))
		newna.Setif(na, ix)
	}
	return NewStrVec(new, newna)
}

// BoolVector implementations ---------------------------------------------------

// Mutate ...
func (v BoolVector) Mutate(f interface{}) Vector {
	fun, ok := f.(func(bool, bool) (bool, bool))
	if !ok {
		if err, ok := f.(error); ok {
			return NewErrVec(fmt.Errorf("interface `f` contains error: %w", err), v.Type())
		}
		return NewErrVec(fmt.Errorf("%w parameter f, expected: `func(string) string`, got `%T`", utils.ErrParamType, f), v.Type())
	}

	new := make([]bool, v.Size())
	newna := make(Set, v.na.Size())
	var na bool
	for ix, val := range v.data {
		new[ix], na = fun(val, v.na.Get(ix))
		newna.Setif(na, ix)
	}
	return NewBoolVec(new, newna)
}
