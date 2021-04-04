package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

func (df Table) P(n string) vec.Vector {
	return df.Pull(n)
}

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

	if err := df.checkCols(n); err != nil {
		return Table{err: err}
	}

	new := make(map[string]vec.Vector, len(n))

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

// Skip ...
func (df Table) Skip(cols []string) Table {

	if err := df.checkCols(cols); err != nil {
		return Table{err: err}
	}

	var empty struct{}
	mcols := make(map[string]struct{}, lim(true, 1, df.size[1]-len(cols)))

	for _, col := range cols {
		mcols[col] = empty
	}

	new := make(map[string]vec.Vector, len(mcols))
	newnames := make([]string, 0, len(mcols))

	for _, val := range df.names {
		if _, ok := mcols[val]; !ok {
			new[val] = df.data[val]
			newnames = append(newnames, val)
		}
	}

	return Table{
		data:  new,
		names: newnames,
		size:  [2]int{df.size[0], len(newnames)},
	}
}

// Skipn ...
func (df Table) Skipn(p []int) Table {
	names := make([]string, len(p))
	for ix, val := range p {
		names[ix] = df.names[val]
	}
	return df.Skip(names)
}

// Cbind ...
// todo test if affects original table!
func (left Table) Cbind(right Table) Table {

	if left.size[0] != right.size[0] {
		return Table{err: fmt.Errorf("table sizes do not match, left rows: %v, right rows: %v", left.size[0], right.size[0])}
	}

	for _, rname := range right.names {
		if _, ok := left.data[rname]; ok {
			left = left.Assign(rname+"_right_", right.data[rname])
		} else {
			left = left.Assign(rname, right.data[rname])
		}
	}
	return left
}
