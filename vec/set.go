package vec

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
func (s Set) Set(i int) Set {
	var empty struct{}
	s[i] = empty
	return s
}

// Get finds whether an element is present in the Set.
func (s Set) Get(i int) bool {
	_, ok := s[i]
	return ok
}

// Del removes the element from the Set.
func (s Set) Del(i int) Set {
	delete(s, i)
	return s
}
