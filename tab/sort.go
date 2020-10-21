package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

func (df Table) Sort(col []string) Table {

	if err := df.checkCols(col); err != nil {
		return Table{err: err}
	}

	ix := df.data[col[0]].Order()
	fmt.Println(ix)

	for c := 0; c < len(col)-1; c++ {
		switch v := df.data[col[c]].(type) {
		case vec.IntVector:
			vals, _ := v.Get()
			start := 0
			check := vals[ix[0]]
			for i := 1; i < len(ix); i++ {
				if vals[ix[i]] != check {
					if i-start == 1 {
						check = vals[ix[i]]
						start = i
						continue
					}
					newIx := df.data[col[c+1]].Loc(ix[start:i]).Order()
					copyIx := make([]int, len(newIx))
					copy(copyIx, ix[start:i])
					for j, pos := range newIx {
						ix[start+j] = copyIx[pos]
					}
					check = vals[ix[i]]
					start = i
					continue
				} else if i == len(ix)-1 {
					newIx := df.data[col[c+1]].Loc(ix[start:]).Order()
					copyIx := make([]int, len(newIx))
					copy(copyIx, ix[start:])
					for j, pos := range newIx {
						ix[start+pos] = copyIx[j]
					}
				}
			}
		case vec.StrVector:
			vals, _ := v.Get()

			start := 0
			check := vals[ix[start]]
			for i := 1; i < len(ix); i++ {
				if vals[ix[i]] != check {
					if i-start == 1 {
						check = vals[ix[i]]
						start = i
						continue
					}
					newIx := df.data[col[c+1]].Loc(ix[start:i]).Order()
					copyIx := make([]int, len(newIx))
					copy(copyIx, ix[start:i])
					for j, pos := range newIx {
						ix[start+j] = copyIx[pos]
					}
					check = vals[ix[i]]
					start = i
				} else if i == len(ix)-1 {
					newIx := df.data[col[c+1]].Loc(ix[start:]).Order()
					copyIx := make([]int, len(newIx))
					copy(copyIx, ix[start:])
					for j, pos := range newIx {
						ix[start+pos] = copyIx[j]
					}
				}
			}
		}
	}
	return df.Rows(ix)
}
