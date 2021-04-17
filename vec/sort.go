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

// IntVector implementations ---------------------------------------------------

// Sort ...
//TODO: NA handling
func (v IntVector) Sort() Vector {
	new := make([]int, v.Size())
	copy(new, v.data)
	sort.Ints(new)
	v.data = new
	return v
}

// Order ...
func (v IntVector) Order() []int {
	return order(sort.IntSlice(v.data), v.na)
}

// StrVector implementations ---------------------------------------------------

// Sort ...
//TODO: NA handling
func (v StrVector) Sort() Vector {
	new := make([]string, v.Size())
	copy(new, v.data)
	sort.Strings(new)
	v.data = new
	return v
}

// Order ...
func (v StrVector) Order() []int {
	return order(sort.StringSlice(v.data), v.na)
}

// BoolVector implementations ---------------------------------------------------

type boolOrdering struct {
	s  BoolVector
	ix []int
}

func lessBool(i, j bool) bool {
	if i == j {
		return false
	}
	if i {
		return j
	}
	return i
}

func (o boolOrdering) Less(i, j int) bool {
	return lessBool(o.s.data[i], o.s.data[j]) || (o.s.na.Get(i) && !o.s.na.Get(j))
}

func (o boolOrdering) Len() int { return o.s.Size() }

func (o boolOrdering) Swap(i, j int) {
	nai := o.s.na.Get(i)
	naj := o.s.na.Get(j)
	if nai && !naj {
		o.s.na.Del(i)
		o.s.na.Set(j)
	}
	if naj && !nai {
		o.s.na.Del(j)
		o.s.na.Set(i)
	}
	o.ix[i], o.ix[j] = o.ix[j], o.ix[i]
}

// Sort ...
//TODO: NA handling
func (v BoolVector) Sort() Vector {
	o := boolOrdering{s: v.Copy().Bool()}
	sort.Stable(o)
	return o.s
}

// Order ...
func (v BoolVector) Order() []int {
	pos := make([]int, v.Size())
	for ix := range pos {
		pos[ix] = ix
	}
	o := boolOrdering{s: v.Copy().Bool(), ix: pos}
	sort.Stable(o)
	return pos

}
