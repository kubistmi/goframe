package tab

type mapping struct {
	col  string
	cols []string
	fun  interface{}
}

type MapFun func() mapping

func MapF(col string, fun interface{}, cols ...string) func() mapping {
	return func() mapping {
		return mapping{
			col, cols, fun,
		}
	}
}

func unwrapMap(maf []MapFun) []mapping {
	mf := make([]mapping, len(maf))
	for ix, val := range maf {
		mf[ix] = val()
	}
	return mf
}
