package tab

import (
	"github.com/kubistmi/goframe/vec"
)

func (df Table) Sort(col []string) Table {

	if err := df.checkCols(col); err != nil {
		return Table{err: err}
	}

	ix := df.data[col[0]].Order()

	for colI, colS := range colN[:len(colN)-1] {
		switch v := df.data[colS].(type) {
		case vec.IntVector:
			vals, _ := v.Get()

			start := 0
			check := vals[ix[start]]
			for i, pos := range ix {

				if vals[pos] != check {
					if i-start == 1 {
						check = vals[pos]
						start = i
						continue
					}
					ixR := df.data[col[colI+1]].Loc(ix[start:i]).Order()
					ixP := make([]int, len(ixR))
					copy(ixP, ix[start:i])
					for a, b := range ixR {
						ix[start+b] = ixP[a]
					}
					check = vals[pos]
					start = i
				} else if i == len(ix)-1 {
					ixR := df.data[col[colI+1]].Loc(ix[start:]).Order()
					ixP := make([]int, len(ixR))
					copy(ixP, ix[start:])
					for a, b := range ixR {
						ix[start+b] = ixP[a]
					}
				}
			}
		case vec.StrVector:
			vals, _ := v.Get()

			check := vals[ix[0]]
			for i, pos := range ix {
				if vals[pos] != check {
					ixR := df.data[col[colI]].Loc(ix[:i]).Order()
					for a, b := range ixR {
						ix[a] = b
					}
					check = vals[pos]
				}
			}
		}
	}
	return df.Rows(ix)
}
