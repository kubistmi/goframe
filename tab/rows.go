package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/utils"
	"github.com/kubistmi/goframe/vec"
)

// Head ...
func (df Table) Head(n int) Table {
	max := lim(false, n, df.size[0])
	p := make([]int, max)
	for i := 0; i < max; i++ {
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

func (df Table) Mask(c []bool) Table {
	if len(c) != df.size[0] {
		return Table{
			err: fmt.Errorf("%w size of slice `c` != size of Table `v`, expected: %v, got: %v", utils.ErrParamVal, df.size[0], len(c)),
		}
	}
	pos := make([]int, 0, df.size[0])
	for ix, val := range c {
		if val {
			pos = append(pos, ix)
		}
	}
	return df.Rows(pos)
}

func (df Table) Fcol(col string, fn interface{}) Table {
	mask, err := df.Pull(col).Check(fn)
	if err != nil {
		return Table{
			err: err,
		}
	}

	return df.Mask(mask)
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
