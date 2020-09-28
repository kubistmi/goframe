package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

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

// Table ... -------------------------------------------------------------------
type Table struct {
	data   []vec.Vector
	names  []string
	inames map[string]int
	index  map[int][]int
	size   [2]int
	err    error
}

// NewDf ...
func NewDf(data map[string]vec.Vector) (Table, error) {

	names := make([]string, 0, len(data))
	new := make([]vec.Vector, 0, len(data))
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
		size:   [2]int{nrow, len(data)},
	}
	return out, nil
}

// Assign ...
func (df Table) Assign(name string, v vec.Vector) Table {
	if v.Size() != df.size[0] {
		return Table{err: fmt.Errorf("wrong vector size, table size: %v, vector size: %v", df.size[0], v.Size())}
	}
	if col, ok := df.inames[name]; ok {
		df.data = append(df.data[:col], df.data[col+1:]...)
		df.names = append(df.names[:col], df.names[col+1:]...)
	}
	df.data = append(df.data, v)
	df.names = append(df.names, name)
	df.inames = inverse(df.names)
	return (df)
}
