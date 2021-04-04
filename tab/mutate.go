package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

// Mutate ...
func (df Table) Mutate(maf ...MapFun) Table {
	mf := unwrapMap(maf)

	for _, val := range mf {
		switch v := df.data[val.cols[0]].(type) {
		case vec.IntVector:
			switch f := val.fun.(type) {
			case func(int, bool) (int, bool):
				if _, ok := df.data[val.col]; !ok {
					df.names = append(df.names, val.col)
				}
				df.data[val.col] = v.Mutate(f)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		case vec.StrVector:
			switch f := val.fun.(type) {
			case func(string, bool) (string, bool):
				if _, ok := df.data[val.col]; !ok {
					df.names = append(df.names, val.col)
				}
				df.data[val.col] = v.Mutate(f)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		}
	}
	return df
}

// MutateM ...
func (df Table) MutateM(maf ...MapFun) Table {
	mf := unwrapMap(maf)

	for ix, val := range mf {
		err := df.checkCols(val.cols)
		return Table{
			err: fmt.Errorf("Error in specification %v: %w", ix, err),
		}
	}

	out := make(map[string]vec.Vector, len(mf))
	for _, val := range mf {
		switch f := val.fun.(type) {
		case func(map[string]string) string:
			res := make([]string, df.size[0])
			resna := vec.NewNA()
			data := make(map[string][]string, len(val.cols))
			nas := vec.NewNA()
			na := vec.NewNA()
			var err error
			parm := make(map[string]string, len(val.cols))
			for _, col := range val.cols {
				data[col], na, err = df.Pull(col).Str().Get()
				if err != nil {
					return (df.setError(fmt.Errorf("error accessign the data")))
				}
				nas.Extend(na)
			}
			for ix := range res {
				if ok := nas.Get(ix); ok {
					resna.Set(ix)
					continue
				}
				for _, col := range val.cols {
					parm[col] = data[col][ix]
				}
				res[ix] = f(parm)
			}
			out[val.col] = vec.NewVec(res, resna)
		case func(map[string]int) int:
			res := make([]int, df.size[0])
			resna := vec.NewNA()
			data := make(map[string][]int, len(val.cols))
			nas := vec.NewNA()
			na := vec.NewNA()
			var err error
			parm := make(map[string]int, len(val.cols))
			for _, col := range val.cols {
				data[col], na, err = df.Pull(col).Int().Get()
				if err != nil {
					return (df.setError(fmt.Errorf("error accessign the data")))
				}
				nas.Extend(na)
			}
			for ix := range res {
				if ok := nas.Get(ix); ok {
					resna.Set(ix)
					continue
				}
				for _, col := range val.cols {
					parm[col] = data[col][ix]
				}
				res[ix] = f(parm)
			}
			out[val.col] = vec.NewVec(res, resna)
		}
	}

	for _, val := range mf {
		df = df.Assign(val.col, out[val.col])
	}

	return df

}

// Pipe ...
func (df Table) Pipe(f func(Table) Table) Table {
	return f(df)
}
