package vec

type NA interface {
	Get(int) bool
	Set(int)
	Setif(bool, int)
	Del(int)
	Size() int
	Extend(NA)
	Collect() []int
	CopyNA() NA
}

func NewNA(size int) NA {
	return NewSet(size)
}

func NewSet(size int) Set {
	if size < 0 {
		return make(Set)
	}
	return make(Set, size)
}

// Set is an alias for map[int]struct{}, it is used to represent the NA values,
// by keeping the positions of the missing data in the Vector.
type Set map[int]struct{}

// Copy allocates a new Set and adds each element of the original into the new one.
func (s Set) Copy() Set {
	new := make(Set)
	for key, val := range s {
		new[key] = val
	}
	return new
}

// Set adds an element into the Set, does nothing if the element exists.
func (s Set) Set(i int) {
	var empty struct{}
	s[i] = empty
}

func (s Set) Setif(b bool, i int) {
	if b {
		s.Set(i)
	}
}

// Get finds whether an element is present in the Set.
func (s Set) Get(i int) bool {
	_, ok := s[i]
	return ok
}

// Del removes the element from the Set.
func (s Set) Del(i int) {
	delete(s, i)
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Extend(n NA) {
	nas := n.Collect()
	for _, val := range nas {
		s.Set(val)
	}
}

func (s Set) Collect() []int {
	out := make([]int, 0, len(s))
	for ix := range s {
		out = append(out, ix)
	}
	return out
}

// Copy implements the copy functionality for NA interface{}
func (s Set) CopyNA() NA {
	return s.Copy()
}
