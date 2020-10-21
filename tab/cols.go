package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

// Pull ...
func (df Table) Pull(n string) vec.Vector {
	if err := df.checkCols([]string{n}); err != nil {
		return vec.NewErrVec(err)
	}
	return df.data[n]
}

// Pulln ...
func (df Table) Pulln(p int) vec.Vector {
	if p >= df.size[1] {
		//? should be own type
		return vec.NewErrVec(fmt.Errorf("wrong position, maximum allowed: %v, got %v", df.size[1]-1, p))
	}
	return df.Pull(df.names[p])
}

// Cols ...
func (df Table) Cols(n []string) Table {

	new := make(map[string]vec.Vector, len(n))

	if err := df.checkCols(n); err != nil {
		return Table{err: err}
	}

	for _, val := range n {
		new[val] = df.data[val]
	}

	return Table{
		data:  new,
		names: n,
		size:  [2]int{df.size[0], len(n)},
	}
}

// Colsn ...
func (df Table) Colsn(p []int) Table {

	n := make([]string, len(p))
	for ix, val := range p {
		if val >= df.size[1] {
			return Table{
				err: fmt.Errorf("wrong position, maximum allowed: %v, got %v", df.size[1]-1, p),
			}
		}
		n[ix] = df.names[val]
	}
	return df.Cols(n)
}
