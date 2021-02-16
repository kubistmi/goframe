package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

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
// Only AND at the moment
func (df Table) Filter(mf map[string]interface{}) Table {

	mask := make([][]bool, 0, len(mf))
	index := make([]int, 0, df.size[0])

	for col, fun := range mf {

		switch v := df.data[col].(type) {
		case vec.IntVector:
			switch f := fun.(type) {
			case func(int) bool:
				val, err := v.Check(f)
				if err != nil {
					return Table{err: fmt.Errorf("error in Check() method in columns %s : %w", col, err)}

				}
				mask = append(mask, val)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) bool, got %T", f),
				}
			}
		case vec.StrVector:
			switch f := fun.(type) {
			case func(string) bool:
				val, err := v.Check(f)
				if err != nil {
					return Table{err: fmt.Errorf("error in Check() method in columns %s : %w", col, err)}

				}
				mask = append(mask, val)
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
