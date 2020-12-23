package vec

// Hash ...
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

	v.hash = nhix
	return v
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

	v.hash = nhix
	return v
}

// GetHash ...
func (v IntVector) GetHash(l int) int {
	return v.hash.lookup[l]
}

// GetHash ...
func (v StrVector) GetHash(l string) int {
	return v.hash.lookup[l]
}

// IsHashed ...
func (v IntVector) IsHashed() bool {
	return len(v.hash.lookup) > 0
}

// IsHashed ...
func (v StrVector) IsHashed() bool {
	return len(v.hash.lookup) > 0
}

// GetHashVals ...
func (v IntVector) GetHashVals() ([]int, int) {
	out := make([]int, v.size)
	for ix, val := range v.obs {
		out[ix] = v.hash.lookup[val]
	}
	return out, v.hash.size
}

// GetHashVals ...
func (v StrVector) GetHashVals() ([]int, int) {
	out := make([]int, v.size)
	for ix, val := range v.obs {
		out[ix] = v.hash.lookup[val]
	}
	return out, v.hash.size
}

// Let's leave this interface her - dunno why it doesnt work
// // Hashable ...
// type Hashable interface {
// 	Vector
// 	GetHashVals() ([]int, int)
// }

// Hash ...
// func Hash(v Vector) Hashable {

// 	switch v := v.(type) {
// 	case IntVector:
// 		nhash := make(map[int]int)

// 		i := 0
// 		for _, val := range v.obs {
// 			if _, ok := nhash[val]; !ok {
// 				nhash[val] = i
// 				i++
// 			}
// 		}

// 		fmt.Println(nhash)

// 		nhix := struct {
// 			lookup map[int]int
// 			size   int
// 		}{nhash, i}

// 		v.hash = nhix
// 		return v

// 	case StrVector:
// 		nhash := make(map[string]int)

// 		i := 0
// 		for _, val := range v.obs {
// 			if _, ok := nhash[val]; !ok {
// 				nhash[val] = i
// 				i++
// 			}
// 		}

// 		fmt.Println(nhash)

// 		nhix := struct {
// 			lookup map[string]int
// 			size   int
// 		}{nhash, i}

// 		v.hash = nhix
// 		return v
// 	}

// 	return StrVector{err: fmt.Errorf("What?")}
// }
