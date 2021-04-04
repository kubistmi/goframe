package vec

import (
	"sort"
)

type ordering struct {
	s  sort.Interface
	ix []int
	na NA
}

func (o ordering) Less(i, j int) bool {
	return o.s.Less(o.ix[i], o.ix[j]) || (o.na.Get(i) && !o.na.Get(j))
}
func (o ordering) Len() int { return o.s.Len() }

func (o ordering) Swap(i, j int) {
	nai := o.na.Get(i)
	naj := o.na.Get(j)
	if nai && !naj {
		o.na.Del(i)
		o.na.Set(j)
	}
	if naj && !nai {
		o.na.Del(j)
		o.na.Set(i)
	}
	o.ix[i], o.ix[j] = o.ix[j], o.ix[i]
}

func order(s sort.Interface, na NA) []int {
	pos := make([]int, s.Len())
	for ix := range pos {
		pos[ix] = ix
	}
	sort.Stable(ordering{s: s, ix: pos, na: na})
	return pos
}

// Sort ...
func (v IntVector) Sort() Vector {
	new := make([]int, v.Size())
	copy(new, v.data)
	sort.Ints(new)
	v.data = new
	return v
}

// Sort ...
func (v StrVector) Sort() Vector {
	new := make([]string, v.Size())
	copy(new, v.data)
	sort.Strings(new)
	v.data = new
	return v
}

// Order ...
func (v IntVector) Order() []int {
	return order(sort.IntSlice(v.data), v.na)
}

// Order ...
func (v StrVector) Order() []int {
	return order(sort.StringSlice(v.data), v.na)
}
