package tab

func (left Table) hashFrom(right Table, cols []string) Table {

	for _, col := range cols {
		if ok := right.data[col].IsHashed(); !ok {
			right.data[col] = right.data[col].Hash()
		}
		left.data[col] = left.data[col].SetHash(right.data[col])
	}
	// todo mising values in index?
	return left.Group(cols)
}

type gmap struct {
	hash  int
	lrows []int
	rrows []int
}

type rmap struct {
	hash, lrow, rrow int
}

// LeftJoin ...
func (left Table) LeftJoin(right Table, on []string) Table {

	if err := left.checkCols(on); err != nil {
		return Table{err: err}
	}
	if err := right.checkCols(on); err != nil {
		return Table{err: err}
	}

	right = right.Group(on)
	left = left.hashFrom(right, on)

	gmaps := make([]gmap, 0, len(left.index.grp))
	lelem := 0

	for hash, lrows := range left.index.grp {
		rrows, ok := right.index.grp[hash]
		if ok {
			gmaps = append(gmaps, gmap{hash, lrows, rrows})
			lelem += len(lrows) * len(rrows)
		} else {
			gmaps = append(gmaps, gmap{hash, lrows, nil})
			lelem += len(lrows)
		}
	}

	hvals := make([]int, 0, lelem)
	lrows := make([]int, 0, lelem)
	rrows := make([]int, 0, lelem)
	for _, val := range gmaps {
		for _, lrow := range val.lrows {
			for _, rrow := range val.rrows {
				hvals = append(hvals, left.index.grp[val.hash][0])
				lrows = append(lrows, lrow)
				rrows = append(rrows, rrow)
			}
		}
	}

	newDf := left.Cols(on).Rows(hvals)
	left = left.Skip(on).Rows(lrows)
	right = right.Skip(on).Rows(rrows)

	return newDf.Cbind(left).Cbind(right)
}
