package vec

// IntVector ... ---------------------------------------------------------------
type IntVector struct {
	obs    []int
	na     Set
	index  []int
	hashix struct {
		lookup map[int]int
		size   int
	}
	size int
	err  error
}

// Size ...
func (v IntVector) Size() int {
	return v.size
}

// GetI ...
func (v IntVector) GetI() (interface{}, Set) {
	new := make([]int, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// Get ...
func (v IntVector) Get() ([]int, Set) {
	new := make([]int, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// Err  ...
func (v IntVector) Err() error {
	return v.err
}

// Copy ...
func (v IntVector) Copy() Vector {

	new := make([]int, v.size)
	copy(new, v.obs)

	return IntVector{
		obs:   new,
		na:    v.na.Copy(),
		size:  v.size,
		index: v.index,
		err:   v.err,
	}
}

// StrVector ... ---------------------------------------------------------------
type StrVector struct {
	obs    []string
	na     Set
	index  []int
	hashix struct {
		lookup map[string]int
		size   int
	}
	size    int
	inverse map[string][]int //?inverse index
	err     error
}

// Size ...
func (v StrVector) Size() int {
	return v.size
}

// GetI ...
func (v StrVector) GetI() (interface{}, Set) {
	new := make([]string, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// Get ...
func (v StrVector) Get() ([]string, Set) {
	new := make([]string, v.Size())
	copy(new, v.obs)
	return new, v.na.Copy()
}

// Err  ...
func (v StrVector) Err() error {
	return v.err
}

// Copy ...
func (v StrVector) Copy() Vector {

	new := make([]string, v.size)
	copy(new, v.obs)

	return StrVector{
		obs:   new,
		na:    v.na.Copy(),
		size:  v.size,
		index: v.index,
		err:   v.err,
	}
}
