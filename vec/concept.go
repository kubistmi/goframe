package vec

import "fmt"

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
