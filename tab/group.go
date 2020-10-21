package tab

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
