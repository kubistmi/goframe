package vec

// Mutate ...
func (v IntVector) Mutate(f func(v int) int) Vector {
	new := make([]int, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
	}
	return IntVector{
		obs:   new,
		na:    v.na,
		index: v.index,
		size:  v.size,
	}
}

// Mutate ...
func (v StrVector) Mutate(f func(v string) string) Vector {
	new := make([]string, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
	}
	return StrVector{
		obs:   new,
		na:    v.na,
		index: v.index,
		size:  v.size,
	}
}
