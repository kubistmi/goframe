package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

// Pull ...
func (df Table) Pull(n string) vec.Vector {
	return df.Pulln(df.inames[n])
}

// Pulln ...
func (df Table) Pulln(p int) vec.Vector {
	if p >= df.size[1] {
		//? should be own type
		return vec.NewErrVec(fmt.Errorf("wrong position, maximum allowed: %v, got %v", df.size[1]-1, p))
	}

	return df.data[p]
}

// Cols ...
func (df Table) Cols(n []string) Table {
	ind := make([]int, 0, len(n))
	for _, val := range n {
		if i, ok := df.inames[val]; ok {
			ind = append(ind, i)
		} else {
			return Table{err: fmt.Errorf("column '%v' not found in df.names", val)}
		}
	}
	return df.Colsn(ind)
}

// Colsn ...
func (df Table) Colsn(p []int) Table {
	new := make([]vec.Vector, len(p))
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
