package vec

// type Groupable interface {
// 	Vector
// 	Group() Vector
// }

// IntVector implementations ---------------------------------------------------

// Group ...
func (v IntVector) Group() Vector {
	groups := make(map[int][]int)

	for ix, val := range v.data {
		slc := groups[val]
		if len(slc) == cap(slc) {
			new := make([]int, len(slc), cap(slc)+100)
			copy(new, slc)
			slc = new
		}
		groups[val] = append(slc, ix)
	}

	v.index = groups
	return v
}

// StrVector implementations ---------------------------------------------------

// Group ...
func (v StrVector) Group() Vector {
	groups := make(map[string][]int)

	for ix, val := range v.data {
		slc := groups[val]
		if len(slc) == cap(slc) {
			news := make([]int, len(slc), cap(slc)+100)
			copy(news, slc)
			slc = news
		}
		groups[val] = append(slc, ix)
	}

	v.index = groups
	return v
}

// BoolVector implementations ---------------------------------------------------

// Group ...
func (v BoolVector) Group() Vector {
	groups := make(map[bool][]int)

	for ix, val := range v.data {
		slc := groups[val]
		if len(slc) == cap(slc) {
			news := make([]int, len(slc), cap(slc)+100)
			copy(news, slc)
			slc = news
		}
		groups[val] = append(slc, ix)
	}

	v.index = groups
	return v
}
