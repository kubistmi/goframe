package tab

import "github.com/kubistmi/goframe/vec"

// GetIndex ...
func (df Table) GetGroups() map[int][]int {
	return df.index.grp
}

// Group ... TESTING GROUP BY
func (df Table) Group(cols []string) Table {

	if err := df.checkCols(cols); err != nil {
		return Table{err: err}
	}

	// HASH EACH VECTOR
	hashtab := make([][]int, len(cols))
	offset := make([]int, len(cols)+1)
	offset[0] = 1

	for j, col := range cols {
		var off int
		ok := df.data[col].IsHashed()
		if !ok {
			df.data[col] = df.data[col].Hash()
		}

		hashtab[j], off = df.data[col].GetHashVals()
		offset[j+1] = off * offset[j]
	}

	nix := make(map[int][]int)
	for i := 0; i < df.size[0]; i++ {
		var chsum int
		for j := 0; j < len(hashtab); j++ {
			chsum += hashtab[j][i] * offset[j]
		}
		nix[chsum] = append(nix[chsum], i)
	}
	df.index = struct {
		cols []string
		grp  map[int][]int
	}{cols, nix}

	return (df)
}

// Agg ...
func (df Table) Agg(mapf map[string]interface{}) Table {

	cols := make([]string, 0, len(mapf))
	for c := range mapf {
		cols = append(cols, c)
	}
	if err := df.checkCols(cols); err != nil {
		return Table{err: err}
	}

	new := make(map[string]vec.Vector, len(mapf)+len(df.index.cols))

	grps := make([]int, 0, len(df.index.grp))
	ix := make([]int, 0, len(grps))
	for grp, rs := range df.index.grp {
		grps = append(grps, grp)
		ix = append(ix, rs[0])
	}

	dfgrp := df.Cols(df.index.cols).Rows(ix)
	for _, n := range dfgrp.names {
		new[n] = dfgrp.data[n]
	}

	cols = make([]string, len(df.index.cols), len(mapf)+len(df.index.cols))
	copy(cols, df.index.cols)
	for col, fun := range mapf {
		switch f := fun.(type) {
		case func(vec.Vector) int:
			aggval := make([]int, len(grps))
			for i, g := range grps {
				aggval[i] = f(df.data[col].Loc(df.index.grp[g]))
			}
			new[col] = vec.NewVec(aggval)
			cols = append(cols, col)
		case func(...vec.Vector) string:
			aggval := make([]string, len(grps))
			for i, g := range grps {
				aggval[i] = f(df.data[col].Loc(df.index.grp[g]))
			}
			new[col] = vec.NewVec(aggval)
			cols = append(cols, col)
		}
	}

	return Table{
		data:  new,
		names: cols,
		size:  [2]int{len(grps), len(cols)},
	}
}
