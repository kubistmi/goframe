package vec

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
