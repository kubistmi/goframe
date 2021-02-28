package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

// Head ...
func (df Table) Head(n int) Table {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	return df.Rows(p)
}

// Rows ...
func (df Table) Rows(p []int) Table {
	new := make(map[string]vec.Vector, df.size[0])
	for ix, val := range df.data {
		new[ix] = val.Loc(p)
		if new[ix].Err() != nil {
			return Table{
				err: fmt.Errorf("Rows: error in Loc() method in column %s : %w", ix, new[ix].Err()),
			}
		}
	}
	return Table{
		data:  new,
		names: df.names,
		size:  [2]int{len(p), df.size[1]},
	}
}

// Filter ...
// Only OR at the moment
func (df Table) Filter(maf ...MapFun) Table {
	mf := unwrapMap(maf)

	mask := make([][]bool, 0, len(mf))
	index := make([]int, 0, df.size[0])

	for _, val := range mf {

		switch v := df.data[val.col].(type) {
		case vec.IntVector:
			switch f := val.fun.(type) {
			case func(int, bool) bool:
				out, err := v.Check(f)
				if err != nil {
					return Table{err: fmt.Errorf("error in Check() method in columns %s : %w", val.col, err)}

				}
				mask = append(mask, out)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) bool, got %T", f),
				}
			}
		case vec.StrVector:
			switch f := val.fun.(type) {
			case func(string, bool) bool:
				out, err := v.Check(f)
				if err != nil {
					return Table{err: fmt.Errorf("error in Check() method in columns %s : %w", val.col, err)}

				}
				mask = append(mask, out)
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
			y = y || mask[j][i]
		}
		if y {
			index = append(index, i)
		}
	}

	return df.Rows(index)

}
