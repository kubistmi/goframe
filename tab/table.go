package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

// Table ... -------------------------------------------------------------------
type Table struct {
	data  map[string]vec.Vector
	names []string
	index struct {
		cols []string
		grp  map[int][]int
	}
	size [2]int
	err  error
}

func (df Table) Err() error {
	return df.err
}

func (df Table) resetErr() Table {
	df.err = nil
	return df
}

func (df Table) setError(err error) Table {
	df.err = err
	return df
}

// NewDf ...
func NewDf(data map[string]vec.Vector) (Table, error) {

	names := make([]string, 0, len(data))

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
	}

	out := Table{
		data:  data,
		names: names,
		size:  [2]int{nrow, len(data)},
	}
	return out, nil
}

// Assign ...
func (df Table) Assign(name string, v vec.Vector) Table {
	if v.Size() != df.size[0] {
		return Table{err: fmt.Errorf("wrong vector size, table size: %v, vector size: %v", df.size[0], v.Size())}
	}
	df.data[name] = v
	df.names = append(df.names, name)
	return df
}

func (df Table) checkCols(col []string) error {
	for _, val := range col {
		if _, ok := df.data[val]; !ok {
			return fmt.Errorf("Column not found in data: %v", val)
		}
	}
	return nil
}
