package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

type mut struct {
	new, old string
	fun      interface{}
}

// Mutate ...
func (df Table) Mutate(mf ...mut) Table {

	for _, val := range mf {
		switch v := df.data[val.old].(type) {
		case vec.IntVector:
			switch f := val.fun.(type) {
			case func(int) int:
				if _, ok := df.data[val.new]; !ok {
					df.names = append(df.names, val.new)
				}
				df.data[val.new] = v.Mutate(f)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		case vec.StrVector:
			switch f := val.fun.(type) {
			case func(string) string:
				if _, ok := df.data[val.new]; !ok {
					df.names = append(df.names, val.new)
				}
				df.data[val.new] = v.Mutate(f)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		}
	}
	return df
}

type mutS struct {
	new string
	old []string
	fun interface{}
}

// MutateM ...
func (df Table) MutateM(mf ...mutS) Table {

	for ix, val := range mf {
		err := df.checkCols(val.old)
		return Table{
			err: fmt.Errorf("Error in specification %v: %w", ix, err),
		}
	}

	out := make(map[string]vec.Vector, len(mf))
	for _, val := range mf {
		switch f := val.fun.(type) {
		case func(map[string]string) string:
			res := make([]string, df.size[0])
			resna := vec.Set{}
			data := make(map[string][]string, len(val.old))
			nas := vec.Set{}
			na := vec.Set{}
			parm := make(map[string]string, len(val.old))
			for _, col := range val.old {
				data[col], na = df.Pull(col).Str().Get()
				nas = nas.Extend(na)
			}
			for ix := range res {
				if ok := nas.Get(ix); ok {
					resna = resna.Set(ix)
					continue
				}
				for _, col := range val.old {
					parm[col] = data[col][ix]
				}
				res[ix] = f(parm)
			}
			out[val.new] = vec.NewVec(res, resna)
		case func(map[string]int) int:
			res := make([]int, df.size[0])
			resna := vec.Set{}
			data := make(map[string][]int, len(val.old))
			nas := vec.Set{}
			na := vec.Set{}
			parm := make(map[string]int, len(val.old))
			for _, col := range val.old {
				data[col], na = df.Pull(col).Int().Get()
				nas = nas.Extend(na)
			}
			for ix := range res {
				if ok := nas.Get(ix); ok {
					resna = resna.Set(ix)
					continue
				}
				for _, col := range val.old {
					parm[col] = data[col][ix]
				}
				res[ix] = f(parm)
			}
			out[val.new] = vec.NewVec(res, resna)
		}
	}

	return df

}

// Pipe ...
func (df Table) Pipe(f func(Table) Table) Table {
	return f(df)
}
