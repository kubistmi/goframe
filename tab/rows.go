package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

// Rows ...
func (df Table) Rows(p []int) Table {
	new := make([]vec.Vector, df.size[0])
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
		case vec.IntVector:
			switch f := fun.(type) {
			case func(int) bool:
				mask = append(mask, v.Find(f))
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) bool, got %T", f),
				}
			}
		case vec.StrVector:
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
