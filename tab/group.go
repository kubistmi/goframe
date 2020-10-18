package tab

import "github.com/kubistmi/goframe/vec"

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
		colH, ok := df.data[col].(vec.Hashable)
		if !ok {
			colH = vec.Hash(df.data[col])
			df.data[col] = colH
		}
		var off int

		hashtab[j], off = colH.GetHashVals()
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
