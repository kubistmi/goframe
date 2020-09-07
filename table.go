package main

import (
	"fmt"
)

// Table ... -------------------------------------------------------------------
type Table struct {
	data   []Vector
	names  []string
	inames map[string]int
	index  []int
	size   [2]int
	err    error
}

// NewDf ...
func NewDf(data map[string]Vector) (Table, error) {

	names := make([]string, 0, len(data))
	new := make([]Vector, 0, len(data))
	// check dimensions
	var nrow int
	for _, val := range data {
		nrow = val.Size()
		break
	}

	for ix, val := range data {
		if val.Size() != nrow {
			return Table{data: nil}, fmt.Errorf("incorrect dimensions in column '%v'", ix)
		}
		names = append(names, ix)
		new = append(new, val)
	}

	out := Table{
		data:   new,
		names:  names,
		inames: inverse(names),
		index:  []int{},
		size:   [2]int{nrow, len(data)},
	}
	return out, nil
}

// Pull ...
func (df Table) Pull(n string) Vector {
	return df.Pulln(df.inames[n])
}

// Pulln ...
func (df Table) Pulln(p int) Vector {
	if p >= df.size[1] {
		//? should be own type
		return StrVector{
			err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", df.size[1]-1, p),
		}
	}
	return df.data[p]
}

// Cols ...
func (df Table) Cols(n []string) Table {
	ind := make([]int, 0, len(n))
	for _, val := range n {
		if i, ok := df.inames[val]; ok {
			ind = append(ind, i)
		} else {
			return Table{err: fmt.Errorf("column '%v' not found in df.names", val)}
		}
	}
	return df.Colsn(ind)
}

// Colsn ...
func (df Table) Colsn(p []int) Table {
	new := make([]Vector, len(p))
	names := make([]string, len(p))
	for ix, val := range p {
		if val >= df.size[1] {
			return Table{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", df.size[1]-1, p),
			}
		}
		//? should be a copy
		new[ix] = df.data[val]
		names[ix] = df.names[val]
	}
	return Table{
		data:   new,
		names:  names,
		inames: inverse(names),
		index:  df.index,
		size:   [2]int{df.size[0], len(p)},
	}
}

// Rows ...
func (df Table) Rows(p []int) Table {
	new := make([]Vector, df.size[0])
	for ix, val := range df.data {
		new[ix] = val.Loc(p)
		if new[ix].Err() != nil {
			return Table{
				err: fmt.Errorf("Rows: error in Loc() method in column %s : %w", df.names[ix], new[ix].Err()),
			}
		}
	}
	return Table{
		data:   new,
		names:  df.names,
		inames: df.inames,
		size:   [2]int{len(p), df.size[1]},
	}
}

// Filter ...
// Only AND at the moment
func (df Table) Filter(mf map[string]interface{}) Table {
	index := make([]int, 0, df.size[0])

	inam := make(map[string]int)
	for n := range mf {
		inam[n] = df.inames[n]
	}

	mask := make([][]bool, 0, len(mf))

	for col, fun := range mf {
		ix := inam[col]

		switch v := df.data[ix].(type) {
		case IntVector:
			switch f := fun.(type) {
			case func(int) bool:
				mask = append(mask, v.Find(f))
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) bool, got %T", f),
				}
			}
		case StrVector:
			switch f := fun.(type) {
			case func(string) bool:
				mask = append(mask, v.Find(f))
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) bool, got %T", f),
				}
			}
		}

	}

	for i := 0; i < df.size[0]; i++ {
		y := true
		for j := 0; j < len(mf); j++ {
			y = y && mask[j][i]
		}
		if y {
			index = append(index, i)
		}
	}

	return df.Rows(index)

}

// Mutate1 ...
func (df Table) Mutate1(mf map[string]interface{}) Table {

	new := make([]Vector, 0, df.size[1])

	for ix, col := range df.names {
		if fun, ok := mf[col]; ok {
			switch v := df.data[ix].(type) {
			case IntVector:
				switch f := fun.(type) {
				case func(int) int:
					new = append(new, v.Mutate(f))
				default:
					return Table{
						err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
					}
				}
			case StrVector:
				switch f := fun.(type) {
				case func(string) string:
					new = append(new, v.Mutate(f))
				default:
					return Table{
						err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
					}
				}
			}

		} else {
			new = append(new, df.data[ix])
		}
	}

	return Table{
		data:   new,
		names:  df.names,
		inames: df.inames,
		index:  df.index,
		size:   df.size,
	}
}

// Mutate ///
func (df Table) Mutate(mf ...mut) Table {

	new := make([]Vector, 0, df.size[1])
	names := make([]string, 0, len(df.names))

	mutm := make(map[string][]mut, len(mf))
	for _, val := range mf {
		mutm[val.old] = append(mutm[val.old], val)
	}

	for ix, col := range df.names {
		if muts, ok := mutm[col]; ok {
			for _, mut := range muts {
				switch v := df.data[ix].(type) {
				case IntVector:
					switch f := mut.fun.(type) {
					case func(int) int:
						new = append(new, v.Mutate(f))
						names = append(names, mut.new)
					default:
						return Table{
							err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
						}
					}
				case StrVector:
					switch f := mut.fun.(type) {
					case func(string) string:
						new = append(new, v.Mutate(f))
						names = append(names, mut.new)
					default:
						return Table{
							err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
						}
					}
				}
			}

		} else {
			new = append(new, df.data[ix])
			names = append(names, df.names[ix])
		}
	}

	return Table{
		data:   new,
		names:  names,
		inames: inverse(names),
		index:  df.index,
		size:   df.size,
	}
}

// Assign ...
func (df Table) Assign(name string, v Vector) Table {
	if v.Size() != df.size[0] {
		return Table{err: fmt.Errorf("wrong vector size, table size: %v, vector size: %v", df.size[0], v.Size())}
	}
	if col, ok := df.inames[name]; ok {
		df.data = append(df.data[:col], df.data[col+1:]...)
		df.names = append(df.names[:col], df.names[col+1:]...)
	}
	df.data = append(df.data, v)
	df.names = append(df.names, name)
	df.inames = inverse(df.names)
	return (df)
}
