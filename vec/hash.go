package vec

// Hash ... TEST HASH INDEX ---------------------------------------------------
func (v IntVector) Hash() Vector {
	nhash := make(map[int]int)

	i := 0
	for _, val := range v.obs {
		if _, ok := nhash[val]; !ok {
			nhash[val] = i
			i++
		}
	}

	nhix := struct {
		lookup map[int]int
		size   int
	}{nhash, i}

	v.hashix = nhix
	return (v)
}

// Hash ...
func (v StrVector) Hash() Vector {
	nhash := make(map[string]int)

	i := 0
	for _, val := range v.obs {
		if _, ok := nhash[val]; !ok {
			nhash[val] = i
			i++
		}
	}

	nhix := struct {
		lookup map[string]int
		size   int
	}{nhash, i}

	v.hashix = nhix
	return (v)
}

// GetHash ...
func (v IntVector) GetHash(l int) int {
	return v.hashix.lookup[l]
}

// GetHash ...
func (v StrVector) GetHash(l string) int {
	return v.hashix.lookup[l]
}

// IsHashed ...
func (v IntVector) IsHashed() bool {
	return len(v.hashix.lookup) > 0
}

// IsHashed ...
func (v StrVector) IsHashed() bool {
	return len(v.hashix.lookup) > 0
}

// GetHashVals ...
func (v IntVector) GetHashVals() ([]int, int) {
	out := make([]int, v.size)
	for ix, val := range v.obs {
		out[ix] = v.hashix.lookup[val]
	}
	return out, v.hashix.size
}

// GetHashVals ...
func (v StrVector) GetHashVals() ([]int, int) {
	out := make([]int, v.size)
	for ix, val := range v.obs {
		out[ix] = v.hashix.lookup[val]
	}
	return out, v.hashix.size
}
