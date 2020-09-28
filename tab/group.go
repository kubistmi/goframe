package tab

// GetIndex ...
func (df Table) GetIndex() map[int][]int {
	return df.index
}

// Group ... TESTING GROUP BY
func (df Table) Group(cols []string) Table {
	pos := make([]int, len(cols))

	//TODO IMPLEMENT CHECKING FOR COL NAMES
	for ix, val := range cols {
		pos[ix] = df.inames[val]
	}

	// HASH EACH VECTOR
	hashtab := make([][]int, len(cols))
	offset := make([]int, len(cols)+1)
	offset[0] = 1

	for j, col := range pos {
		if !df.data[col].IsHashed() {
			df.data[col] = df.data[col].Hash()
		}
		var off int

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
	df.index = nix
	return (df)

}
