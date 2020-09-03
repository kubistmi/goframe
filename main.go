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
	Loc(p int) Vector
}

// IntVector ...
type IntVector struct {
	obs   []int
	name  string
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

// StrVector ...
type StrVector struct {
	obs     []string
	name    string
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

// Loc ...
func (v IntVector) Loc(p int) Vector {
	if (p + 1) > v.size {
		return IntVector{
			err: fmt.Errorf("wrong position, vector size: %v, got %v", v.size, p),
		}
	}
	return IntVector{
		obs:  []int{v.obs[p]},
		name: v.name,
		size: 1,
	}
}

// Loc ...
func (v StrVector) Loc(p int) Vector {
	if (p + 1) > v.size {
		return StrVector{
			err: fmt.Errorf("wrong position, vector size: %v, got %v", v.size, p),
		}
	}
	return StrVector{
		obs:  []string{v.obs[p]},
		name: v.name,
		size: 1,
	}
}

// Table ...
type Table struct {
	data  []Vector
	names []string
	index []int
	size  [2]int
	err   error
}

// Loc ...
func (df Table) Loc(p int) Vector {
	if (p + 1) > df.size[1] {
		// should be own type?
		return IntVector{
			err: fmt.Errorf("wrong position, table size: %v, got %v", df.size, p),
		}
	}
	return df.data[p]
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
			return Table{data: nil}, fmt.Errorf("incorrect dimensions in column %v", ix)
		}
		names = append(names, ix)
		new = append(new, val)
	}

	out := Table{
		data:  new,
		names: names,
		index: []int{},
		size:  [2]int{nrow, len(data)},
	}
	return out, nil
}

// Help constructing slices?
func c(p ...int) []int {
	return p
}

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
	for ix, val := range df.Loc(0).Loc(5).Get().([]int) {
		fmt.Printf("%v = %v\n", ix, val)
	}
}
