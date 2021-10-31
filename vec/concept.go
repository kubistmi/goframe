package vec

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
)

type Elem struct {
	val interface{}
	na  bool
	err error
}

func (e Elem) IsNa() bool {
	return e.na
}

func (e Elem) Err() error {
	return e.err
}

func (e Elem) Val() interface{} {
	return e.val
}

// Should this be using Vector.Elem()?
func (v StrVector) GetElem(i int) Elem {
	if i >= v.Size() || i < 0 {
		return Elem{err: fmt.Errorf("%w position i is higher than size of Vector, expected: %v, got: %v", utils.ErrParamVal, v.Size(), i)}
	}
	return Elem{v.data[i], v.na.Get(i), nil}
}

func (v IntVector) GetElem(i int) Elem {
	if i >= v.Size() || i < 0 {
		return Elem{err: fmt.Errorf("%w position i is higher than size of Vector, expected: %v, got: %v", utils.ErrParamVal, v.Size(), i)}
	}
	return Elem{v.data[i], v.na.Get(i), nil}
}

func (v BoolVector) GetElem(i int) Elem {
	if i >= v.Size() || i < 0 {
		return Elem{err: fmt.Errorf("%w position i is higher than size of Vector, expected: %v, got: %v", utils.ErrParamVal, v.Size(), i)}
	}
	return Elem{v.data[i], v.na.Get(i), nil}
}

func (v StrVector) ElemI(i int) (interface{}, bool) {
	if i >= v.Size() {
		var e interface{}
		return e, true
	}
	return v.data[i], v.na.Get(i)
}

func (v IntVector) ElemI(i int) (interface{}, bool) {
	if i >= v.Size() {
		var e interface{}
		return e, true
	}
	return v.data[i], v.na.Get(i)
}

func (v BoolVector) ElemI(i int) (interface{}, bool) {
	if i >= v.Size() {
		var e interface{}
		return e, true
	}
	return v.data[i], v.na.Get(i)
}

func (v StrVector) Elem(i int) (string, bool) {
	if i >= v.Size() {
		var e string
		return e, true
	}
	return v.data[i], v.na.Get(i)
}

func (v IntVector) Elem(i int) (int, bool) {
	if i >= v.Size() {
		var e int
		return e, true
	}
	return v.data[i], v.na.Get(i)
}

func (v BoolVector) Elem(i int) (bool, bool) {
	if i >= v.Size() {
		var e bool
		return e, true
	}
	return v.data[i], v.na.Get(i)
}

func (v StrVector) Is(val interface{}) func() (bool, error) {
	vt, ok := val.(string)
	if !ok {
		return func() (bool, error) {
			return false, fmt.Errorf("wrong value supplied")
		}
	}

	i := 0
	size := v.Size()

	return func() (bool, error) {
		if i >= size {
			return false, fmt.Errorf("iterator exhausted")
		}
		i++
		return v.data[i-1] == vt, nil
	}
}

func (v IntVector) Is(val interface{}) func() (bool, error) {

	vt, ok := val.(int)
	if !ok {
		return func() (bool, error) {
			return false, fmt.Errorf("wrong value supplied")
		}
	}

	i := 0
	size := v.Size()

	return func() (bool, error) {
		if i >= size {
			return false, fmt.Errorf("iterator exhausted")
		}
		i++
		return v.data[i-1] == vt, nil
	}
}

func (v StrVector) Iter() (func() (string, bool), error) {

	if v.err != nil {
		return func() (string, bool) {
			return "", false
		}, v.err
	}

	i := 0
	size := v.Size()

	return func() (string, bool) {
		if i >= size {
			return "", true
		}
		i++
		return v.data[i-1], v.na.Get(i - 1)
	}, nil
}

func (v IntVector) Iter() (func() (int, bool), error) {

	if v.err != nil {
		return func() (int, bool) {
			return 0, false
		}, v.err
	}
	i := 0
	size := v.Size()

	return func() (int, bool) {
		if i >= size {
			return 0, true
		}
		i++
		return v.data[i-1], v.na.Get(i - 1)
	}, nil
}
