package vec

import (
	"sort"
)

type ordering struct {
	s  sort.Interface
	ix []int
}

func (o ordering) Less(i, j int) bool { return o.s.Less(o.ix[i], o.ix[j]) }
func (o ordering) Len() int           { return o.s.Len() }
func (o ordering) Swap(i, j int)      { o.ix[i], o.ix[j] = o.ix[j], o.ix[i] }

func order(s sort.Interface) []int {
	pos := make([]int, s.Len())
	for ix := range pos {
		pos[ix] = ix
	}
	sort.Stable(ordering{s: s, ix: pos})
	return pos
}

// Sort ...
func (v IntVector) Sort() Vector {
	new := make([]int, v.Size())
	copy(new, v.obs)
	sort.Ints(new)
	v.obs = new
	return v
}

// Sort ...
func (v StrVector) Sort() Vector {
	new := make([]string, v.Size())
	copy(new, v.obs)
	sort.Strings(new)
	v.obs = new
	return v
}

// Order ...
func (v IntVector) Order() []int {
	return order(sort.IntSlice(v.obs))
}

// Order ...
func (v StrVector) Order() []int {
	return order(sort.StringSlice(v.obs))
}
