package tab

import "github.com/kubistmi/goframe/vec"

func (df Table) Gen(fun interface{}) Table {

	switch fO := fun.(type) {
	case func() func() (int, bool):
		f := fO()
		new := make([]int, df.size[0])
		na := make(vec.Set)
		for ix := range new {
			v, n := f()
			if n {
				na.Set(ix)
			} else {
				new[ix] = v
			}
		}
		df = df.Assign("new", vec.NewVec(new, na))
	case func() func() (string, bool):
		f := fO()
		new := make([]string, df.size[0])
		na := make(vec.Set)
		for ix := range new {
			v, n := f()
			if n {
				na.Set(ix)
			} else {
				new[ix] = v
			}
		}
		df = df.Assign("new", vec.NewVec(new, na))
	case error:
		df.err = fO
	}
	return df
}
