package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

func (df Table) Mut(maf ...MapFun) Table {
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

// MutateMap ...
//TODO: implement NA handling
func (df Table) MutateMap(maf ...MapFun) Table {
	mf := unwrapMap(maf)

	for ix, val := range mf {
		err := df.checkCols(val.cols)
		if err != nil {
			return Table{
				err: fmt.Errorf("Error in specification %v: %w", ix, err),
			}
		}
	}

	out := make(map[string]vec.Vector, len(mf))
	for _, val := range mf {
		switch f := val.fun.(type) {
		case func(map[string]interface{}) string:
			res := make([]string, df.size[0])
			resna := vec.NewNA(-1)
			parm := make(map[string]interface{}, len(val.cols))
			for ix := range res {
				for _, col := range val.cols {
					parm[col], _ = df.P(col).ElemI(ix)
				}
				res[ix] = f(parm)
			}
			out[val.col] = vec.NewVec(res, resna)
		case func(map[string]interface{}) int:
			res := make([]int, df.size[0])
			resna := vec.NewNA(-1)
			parm := make(map[string]interface{}, len(val.cols))
			for ix := range res {
				for _, col := range val.cols {
					parm[col], _ = df.P(col).ElemI(ix)
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
