package tab

import "github.com/kubistmi/goframe/vec"

func (df Table) Spread(cols []string, spreadcol string, valcol string) Table {
	if err := df.checkCols(cols); err != nil {
		return Table{err: err}
	}

	// compare selected vs grouped columns
	//todo make sure this doesn't corrupt grouping of original table!
	//gcols := append(cols, spreadcol)
	gc := df.GroupCols()
	if len(gc) != len(cols) {
		df = df.Group(cols)
	} else {
		gcm := make(map[string]struct{})
		for _, v := range gc {
			gcm[v] = struct{}{}
		}
		same := true
		for _, v := range cols {
			if _, ok := gcm[v]; !ok {
				same = false
				break
			}
		}
		if !same {
			df = df.Group(cols)
		}
	}

	spread := df.Pull(spreadcol).ToStr()
	//todo handle NAs
	vals, _, err := spread.Get()
	if err != nil {
		return Table{err: err}
	}

	newcols := spread.Unique()

	newix := make([]int, len(df.index.grp))
	newdata := make(map[string]vec.Vector, len(newcols))

	switch v := df.Pull(valcol).(type) {
	case vec.StrVector:

		data, nas, err := v.Get()
		if err != nil {
			return Table{err: err}
		}
		d := make(map[string][]string, len(newcols))
		for _, c := range newcols {
			d[c] = make([]string, len(df.index.grp))
		}
		newna := make(map[string]vec.NA, len(newcols))
		for _, c := range newcols {
			newna[c] = vec.NewNA(0)
		}

		ix := 0
		for _, row := range df.index.grp {
			for _, srow := range row {
				newix[ix] = srow
				nc := vals[srow]
				if nas.Get(srow) {
					newna[nc].Set(ix)
				} else {
					d[nc][ix] = data[srow]
				}
			}
			ix++
		}

		for ix := range d {
			newdata[ix] = vec.NewStrVec(d[ix], newna[ix])
		}
	case vec.IntVector:

		data, nas, err := v.Get()
		if err != nil {
			return Table{err: err}
		}
		d := make(map[string][]int, len(newcols))
		for _, c := range newcols {
			d[c] = make([]int, len(df.index.grp))
		}
		newna := make(map[string]vec.NA, len(newcols))
		for _, c := range newcols {
			newna[c] = vec.NewNA(0)
		}

		ix := 0
		for _, row := range df.index.grp {
			for _, srow := range row {
				newix[ix] = srow
				nc := vals[srow]
				if nas.Get(srow) {
					newna[nc].Set(ix)
				} else {
					d[nc][ix] = data[srow]
				}
			}
			ix++
		}

		for ix := range d {
			newdata[ix] = vec.NewIntVec(d[ix], newna[ix])
		}

	case vec.BoolVector:

		data, nas, err := v.Get()
		if err != nil {
			return Table{err: err}
		}
		d := make(map[string][]bool, len(newcols))
		for _, c := range newcols {
			d[c] = make([]bool, len(df.index.grp))
		}
		newna := make(map[string]vec.NA, len(newcols))
		for _, c := range newcols {
			newna[c] = vec.NewNA(0)
		}

		ix := 0
		for _, row := range df.index.grp {
			for _, srow := range row {
				newix[ix] = srow
				nc := vals[srow]
				if nas.Get(srow) {
					newna[nc].Set(ix)
				} else {
					d[nc][ix] = data[srow]
				}
			}
			ix++
		}

		for ix := range d {
			newdata[ix] = vec.NewBoolVec(d[ix], newna[ix])
		}
	}

	//newval := df.Pull(valcol)
	newdf := df.Cols(cols).Rows(newix)
	for ix, val := range newdata {
		newdf = newdf.Assign(ix, val)
	}

	return newdf
}
