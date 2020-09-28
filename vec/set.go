package vec

func cpMap(a map[int]bool) map[int]bool {
	new := make(map[int]bool)
	for key, val := range a {
		new[key] = val
	}
	return new
}

type Set map[int]struct{}

func (s Set) Copy() Set {
	new := make(Set)
	for key, val := range s {
		new[key] = val
	}
	return new
}

func (s Set) Set(i int) Set {
	var empty struct{}
	s[i] = empty
	return s
}

func (s Set) Get(i int) bool {
	_, ok := s[i]
	return ok
}
