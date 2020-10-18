package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

type mapf map[string]interface{}

// Mutate1 ...
func (df Table) Mutate1(mf mapf) Table {

	new := make([]vec.Vector, 0, df.size[1])

	//TODO: DONT COPY!
	// Should be implemented together with Table as map[string]Vector
	for ix, col := range df.names {
		if fun, ok := mf[col]; ok {
			switch v := df.data[ix].(type) {
			case vec.IntVector:
				switch f := fun.(type) {
				case func(int) int:
					new = append(new, v.Mutate(f))
				default:
					return Table{
						err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
					}
				}
			case vec.StrVector:
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

type mut2 struct {
	new string
	old []string
	fun func(...interface{}) interface{}
}

// Mutate2 ...
func (df Table) Mutate2(mf ...mut2) Table {

	new := make([]vec.Vector, 0, len(mf))
	newnames := make([]string, 0, len(mf))

	for _, str := range mf {

		newnames = append(newnames, str.new)

		// GET COLUMN POSITION FROM NAMES
		cols := make([]int, 0, len(str.old))
		for _, col := range str.old {
			cols = append(cols, df.inames[col])
		}

		args := make([]interface{}, len(str.old))

		for ix, col := range cols {
			args[ix], _ = df.data[col].Elem(0) //nas[ix]
		}

		switch str.fun(args...).(type) {
		case int:
			out := make([]int, df.size[0])
			//nas := make([]bool, len(str.old))
			for i := 0; i < df.size[0]; i++ {
				for ix, col := range cols {
					args[ix], _ = df.data[col].Elem(i) //nas[ix]
				}
				out[i] = str.fun(args...).(int)
				new = append(new, vec.NewVec(out))
			}

		case string:
			out := make([]string, df.size[0])
			//nas := make([]bool, len(str.old))
			for i := 0; i < df.size[0]; i++ {
				for ix, col := range cols {
					args[ix], _ = df.data[col].Elem(i) //nas[ix]
				}
				out[i] = str.fun(args...).(string)
				new = append(new, vec.NewVec(out))
			}
		}
	}

	return Table{
		data:   append(df.data, new...),
		names:  append(df.names, newnames...),
		inames: inverse(append(df.names, newnames...)),
		index:  df.index,
		size:   df.size,
	}
}

type mut struct {
	new, old string
	fun      interface{}
}

// Mutate ///
func (df Table) Mutate(mf ...mut) Table {

	new := make([]vec.Vector, 0, df.size[1])
	names := make([]string, 0, len(df.names))
	//TODO: DONT COPY!
	// Should be implemented together with Table as map[string]Vector
	mutm := make(map[string][]mut, len(mf))
	for _, val := range mf {
		mutm[val.old] = append(mutm[val.old], val)
	}

	for ix, col := range df.names {
		if muts, ok := mutm[col]; ok {
			for _, mut := range muts {
				switch v := df.data[ix].(type) {
				case vec.IntVector:
					switch f := mut.fun.(type) {
					case func(int) int:
						new = append(new, v.Mutate(f))
						names = append(names, mut.new)
					default:
						return Table{
							err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
						}
					}
				case vec.StrVector:
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

// Pipe ...
func (df Table) Pipe(f func(Table) Table) Table {
	return f(df)
}
