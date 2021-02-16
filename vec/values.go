package vec

import (
	"fmt"
)

// GetI ...
func (v IntVector) GetI() (interface{}, Set) {
	new := make([]int, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// GetI ...
func (v StrVector) GetI() (interface{}, Set) {
	new := make([]string, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// Get ...
func (v IntVector) Get() ([]int, Set) {
	new := make([]int, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// Get ...
func (v StrVector) Get() ([]string, Set) {
	new := make([]string, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// AssignI ...
//TODO: is it safe?
func (v IntVector) AssignI(p int, val interface{}, na bool) Vector {
	tval, ok := val.(int)
	if !ok {
		v.err = fmt.Errorf("Wrong value type, expected: %T got: %T", v.obs[0], val)
		return v
	}
	if p >= v.Size() {
		v.err = fmt.Errorf("Wrong assignment position, max: %v got: %v", v.Size(), p)
		return v
	}

	v.obs[p] = tval
	if na {
		v.na.Set(p)
	}
	return v
}

// AssignI ...
//TODO: is it safe?
func (v StrVector) AssignI(p int, val interface{}, na bool) Vector {
	tval, ok := val.(string)
	if !ok {
		v.err = fmt.Errorf("Wrong value type, expected: %T got: %T", v.obs[0], val)
		return v
	}
	if p >= v.Size() {
		v.err = fmt.Errorf("Wrong assignment position, max: %v got: %v", v.Size(), p)
		return v
	}

	v.obs[p] = tval
	if na {
		v.na.Set(p)
	}
	return v
}

// AssignM ...
//TODO: is it safe?
func (v StrVector) AssignM(p []int, val interface{}, na Set) Vector {
	tval, ok := val.([]string)
	if !ok {
		v.err = fmt.Errorf("Wrong value type, expected: %T got: %T", v.obs[0], val)
		return v
	}

	for ix, val := range p {
		v = v.Assign(val, tval[ix], na.Get(ix))
	}

	return v
}

// AssignM ...
//TODO: is it safe?
func (v IntVector) AssignM(p []int, val interface{}, na Set) Vector {
	tval, ok := val.([]int)
	if !ok {
		v.err = fmt.Errorf("Wrong value type, expected: %T got: %T", v.obs[0], val)
		return v
	}

	for ix, val := range p {
		v = v.Assign(val, tval[ix], na.Get(ix))
	}

	return v
}

// Assign ...
//TODO: is it safe?
func (v IntVector) Assign(p int, val int, na bool) IntVector {
	if p >= v.Size() {
		v.err = fmt.Errorf("Wrong assignment position, max: %v got: %v", v.Size(), p)
		return v
	}

	v.obs[p] = val
	if na {
		v.na.Set(p)
	}
	return v
}

// Assign ...
//TODO: is it safe?
func (v StrVector) Assign(p int, val string, na bool) StrVector {
	if p >= v.Size() {
		v.err = fmt.Errorf("Wrong assignment position, max: %v got: %v", v.Size(), p)
		return v
	}

	v.obs[p] = val
	if na {
		v.na.Set(p)
	}
	return v
}

// Copy ...
func (v IntVector) Copy() Vector {

	new := make([]int, v.size)
	copy(new, v.obs)

	return IntVector{
		obs:  new,
		na:   v.na.Copy(),
		size: v.size,
		hash: v.hash,
		err:  v.err,
	}
}

// Copy ...
func (v StrVector) Copy() Vector {

	new := make([]string, v.size)
	copy(new, v.obs)

	return StrVector{
		obs:  new,
		na:   v.na.Copy(),
		size: v.size,
		hash: v.hash,
		err:  v.err,
	}
}
