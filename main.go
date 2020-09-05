package main

import (
	"fmt"
	"log"
)

// Vector ...
type Vector interface {
	Size() int
	//Get[T vector_type]() []T
	Get() interface{}
	Loc(p []int) Vector
	Err() error
}

// IntVector ... ---------------------------------------------------------------
type IntVector struct {
	obs   []int
	index []int
	size  int
	err   error
}

// Size ...
func (v IntVector) Size() int {
	return v.size
}

// Get ...
func (v IntVector) Get() interface{} {
	return v.obs
}

// Err  ...
func (v IntVector) Err() error {
	return v.err
}

// Loc ...
func (v IntVector) Loc(p []int) Vector {
	new := make([]int, len(p))
	for ix, val := range p {
		if val >= v.Size() {
			return IntVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
		}
		new[ix] = v.obs[val]
	}
	return IntVector{
		obs:  new,
		size: len(p),
	}
}

// StrVector ... ---------------------------------------------------------------
type StrVector struct {
	obs     []string
	index   []int
	size    int
	inverse map[string][]int
	err     error
}

// Size ...
func (v StrVector) Size() int {
	return v.size
}

// Get ...
func (v StrVector) Get() interface{} {
	return v.obs
}

// Err  ...
func (v StrVector) Err() error {
	return v.err
}

// Loc ...
func (v StrVector) Loc(p []int) Vector {
	new := make([]string, len(p))
	for ix, val := range p {
		if val >= v.Size() {
			return StrVector{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", v.Size()-1, val),
			}
		}
		new[ix] = v.obs[val]
	}
	return StrVector{
		obs: new,
	}
}

// Table ... -------------------------------------------------------------------
type Table struct {
	data   []Vector
	names  []string
	inames map[string]int
	index  []int
	size   [2]int
	err    error
}

// Pull ...
func (df Table) Pull(p int) Vector {
	if p >= df.size[1] {
		//? should be own type
		return StrVector{
			err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", df.size[1]-1, p),
		}
	}
	return df.data[p]
}

// Cols ...
func (df Table) Cols(p []int) Table {
	new := make([]Vector, len(p))
	names := make([]string, len(p))
	for ix, val := range p {
		if val >= df.size[1] {
			return Table{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", df.size[1]-1, p),
			}
		}
		new[ix] = df.data[val]
		names[ix] = df.names[val]
	}
	return Table{
		data:   new,
		names:  names,
		inames: inverse(names),
		index:  df.index,
		size:   [2]int{df.size[0], len(p)},
	}
}

// Rows ...
func (df Table) Rows(p []int) Table {
	new := make([]Vector, df.size[0])
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

// NewVec ...
func NewVec(data interface{}) (Vector, error) {
	switch t := data.(type) {
	case []int:
		return IntVector{
			obs:  t,
			size: len(t),
		}, nil

	case []string:
		return StrVector{
			obs:  t,
			size: len(t),
		}, nil
	default:
		return nil, fmt.Errorf("wrong data type, expected []int or []string, got %T", t)
	}
}

// NewDf ...
func NewDf(data map[string]Vector) (Table, error) {

	names := make([]string, 0, len(data))
	new := make([]Vector, 0, len(data))
	// check dimensions
	var nrow int
	for _, val := range data {
		nrow = val.Size()
		break
	}

	for ix, val := range data {
		if val.Size() != nrow {
			return Table{data: nil}, fmt.Errorf("incorrect dimensions in column '%v'", ix)
		}
		names = append(names, ix)
		new = append(new, val)
	}

	out := Table{
		data:   new,
		names:  names,
		inames: inverse(names),
		index:  []int{},
		size:   [2]int{nrow, len(data)},
	}
	return out, nil
}

// WTH IS THIS? ----------------------------------------------------------------
func (v IntVector) mutate(f func(v int) int) Vector {
	new := make([]int, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
	}
	return IntVector{
		obs:   new,
		index: v.index,
		size:  v.size,
	}
}

func (v StrVector) mutate(f func(v string) string) Vector {
	new := make([]string, v.Size())
	for ix, val := range v.obs {
		new[ix] = f(val)
	}
	return StrVector{
		obs:   new,
		index: v.index,
		size:  v.size,
	}
}

// func (df Table) mutate(mf map[string]func(v Vector) Vector) Table {
// 	for col, fun := range mf {
// 		ix := df.inames[col]
// 		switch t := df.data[ix].(type) {
// 		case IntVector:
// 			t.mutate(fun)

// 		}
// 	}
// }

// Mutate ...
func (df Table) Mutate(mf map[string]interface{}) Table {

	new := make([]Vector, 0, len(mf))
	names := make([]string, 0, len(mf))
	ind := make([]int, 0, df.size[1]-len(mf))
	oldnam := make([]string, 0, df.size[1]-len(mf))

	for col, fun := range mf {
		ix := df.inames[col]

		switch v := df.data[ix].(type) {
		case IntVector:
			switch f := fun.(type) {
			case func(int) int:
				new = append(new, v.mutate(f))
				names = append(names, col)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		case StrVector:
			switch f := fun.(type) {
			case func(string) string:
				new = append(new, v.mutate(f))
				names = append(names, col)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		}
	}

	for ix, val := range df.names {
		if _, ok := mf[val]; !ok {
			ind = append(ind, ix)
			oldnam = append(oldnam, df.names[ix])
		}
	}

	names = append(oldnam, names...)

	return Table{
		data:   append(df.Cols(ind).data, new...),
		names:  names,
		inames: inverse(names),
		index:  []int{},
		size:   df.size,
	}
}

// WTH IS THIS? ----------------------------------------------------------------

func inverse(names []string) map[string]int {
	inames := make(map[string]int)
	for ix, val := range names {
		inames[val] = ix
	}
	return inames
}

// Help constructing slices?
func c(p ...int) []int {
	return p
}

type mapf map[string]interface{}

func main() {
	vecI, err := NewVec([]int{0, 1, 2, 3, 4, 5})
	if err != nil {
		log.Fatal(err)
	}
	vecS, err := NewVec([]string{"a", "b", "c", "d", "e", "f"})
	if err != nil {
		log.Fatal(err)
	}
	df, err := NewDf(map[string]Vector{"ints": vecI, "strs": vecS})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df)
	for ix, val := range df.Pull(0).Loc([]int{5}).Get().([]int) {
		fmt.Printf("%v = %v\n", ix, val)
	}

	fmt.Println(df.Rows(c(1, 2)))
	fmt.Printf("%v\n", df.Rows(c(20)).err)

	b := df.Mutate(mapf{"ints": func(i int) int {
		return i * 3
	}})

	fmt.Println(b)
}
