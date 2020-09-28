package main

import (
	"fmt"
	"reflect"
)

// REFLECT BASED MUTATE --------------------------------------------------------
func (df Table) mutateM(f interface{}) {

	argS := reflect.TypeOf(f).In(0)

	nargs := argS.NumField()
	fields := make([]reflect.StructField, 0, nargs)
	cols := make([]string, 0, nargs)
	types := make([]string, 0, nargs)
	for i := 0; i < nargs; i++ {
		cols = append(cols, argS.Field(i).Name)
		types = append(types, argS.Field(i).Type.Name())
		fields = append(fields, reflect.StructField{Name: argS.Field(i).Name, Type: reflect.SliceOf(argS.Field(i).Type)})
	}

	argv := reflect.New(reflect.StructOf(fields))

	for _, col := range cols {
		ix := df.inames[col]
		switch vec := df.data[ix].(type) {
		case IntVector:
			argv.Elem().FieldByName(col).Set(reflect.ValueOf(vec.obs))
		case StrVector:
			argv.Elem().FieldByName(col).Set(reflect.ValueOf(vec.obs))
		}
	}

	fmt.Println(argv)
	fmt.Println(types)

	// switch fun := f.(type) {
	// case func(a) string:
	// 	new := make([]string, df.size[0])
	// 	for ix, val := range new {
	// 		argv.Elem().Field(0).Int()
	// 	}
	// }

}

// type a struct {
// 	Ints int
// 	Strs string
// }

// df.mutateM(func(p a) string {
// 	var out string
// 	for i := 0; i < p.Ints; i++ {
// 		out = out + p.Strs
// 	}
// 	return out
// })

// MAP BASED TABLE STRUCTURE ---------------------------------------------------

type Table2 struct {
	strs  map[string]StrVector
	ints  map[string]IntVector
	names []string
	index []int
	size  [2]int
	err   error
}

func NewDf2(data map[string]Vector) Table2 {

	names := make([]string, 0, len(data))
	strs := make(map[string]StrVector)
	ints := make(map[string]IntVector)
	// check dimensions
	var nrow int
	for _, val := range data {
		nrow = val.Size()
		break
	}

	for ix, val := range data {
		if val.Size() != nrow {
			return Table2{err: fmt.Errorf("incorrect dimensions in column '%v'", ix)}
		}
		names = append(names, ix)
		switch t := val.(type) {
		case IntVector:
			ints[ix] = t
		case StrVector:
			strs[ix] = t
		}
	}

	out := Table2{
		strs:  strs,
		ints:  ints,
		names: names,
		index: []int{},
		size:  [2]int{nrow, len(data)},
	}
	return out
}

func (df Table2) Filter(mf map[string]interface{}) Table2 {
	index := make([]int, 0, df.size[0])
	mask := make([][]bool, 0, len(mf))

	for col, fun := range mf {
		if vec, ok := df.strs[col]; ok {
			switch f := fun.(type) {
			case func(string) bool:
				mask = append(mask, vec.Find(f))
			default:
				return Table2{
					err: fmt.Errorf("wrong function definition, expected func(int) bool, got %T", f),
				}
			}
		} else if vec, ok := df.ints[col]; ok {
			switch f := fun.(type) {
			case func(int) bool:
				mask = append(mask, vec.Find(f))
			default:
				return Table2{
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

func (df Table2) Rows(p []int) Table2 {
	strs := make(map[string]StrVector, df.size[0])
	ints := make(map[string]IntVector, df.size[0])

	for col, val := range df.strs {
		strs[col] = val.Loc(p).(StrVector)
		if strs[col].Err() != nil {
			return Table2{
				err: fmt.Errorf("Rows: error in Loc() method in column %s : %w", col, strs[col].Err()),
			}
		}
	}

	for col, val := range df.ints {
		ints[col] = val.Loc(p).(IntVector)
		if ints[col].Err() != nil {
			return Table2{
				err: fmt.Errorf("Rows: error in Loc() method in column %s : %w", col, ints[col].Err()),
			}
		}
	}

	return Table2{
		strs:  strs,
		ints:  ints,
		names: df.names,
		size:  [2]int{len(p), df.size[1]},
	}
}

type mut3 struct {
	new string
	old []string
	fun interface{}
}

func (df Table2) Mutate2(mf ...mut3) Table2 {

	// strs := make([]Vector, 0, df.size[1])
	// ints := make([]Vector, 0, df.size[1])
	// names := make([]string, 0, len(df.names))

	intm := make(map[string][]int)
	strm := make(map[string][]string)
	for _, op := range mf {

		for _, oldC := range op.old {
			if col, ok := df.strs[oldC]; ok {
				strm[oldC], _ = col.Get()
			} else if col, ok := df.ints[oldC]; ok {
				intm[oldC], _ = col.Get()
			}
		}
	}

	// should I even continue?
	// if muts, ok := mutm[col]; ok {
	// 	for _, mut := range muts {
	// 		switch v := df.data[ix].(type) {
	// 		case IntVector:
	// 			switch f := mut.fun.(type) {
	// 			case func(int) int:
	// 				new = append(new, v.Mutate(f))
	// 				names = append(names, mut.new)
	// 			default:
	// 				return Table2{
	// 					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
	// 				}
	// 			}
	// 		case StrVector:
	// 			switch f := mut.fun.(type) {
	// 			case func(string) string:
	// 				new = append(new, v.Mutate(f))
	// 				names = append(names, mut.new)
	// 			default:
	// 				return Table2{
	// 					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
	// 				}
	// 			}
	// 		}
	// 	}

	// } else {
	// 	new = append(new, df.data[ix])
	// 	names = append(names, df.names[ix])
	// }

	// return Table2{
	// 	data:   new,
	// 	names:  names,
	// 	inames: inverse(names),
	// 	index:  df.index,
	// 	size:   df.size,
	// }
	return Table2{err: fmt.Errorf("unimplemented")}
}
