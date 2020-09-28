package vec

// Elem ... TEST ACCESS -------------------------------------------------------
func (v IntVector) Elem(i int) (interface{}, bool) {
	return v.obs[i], v.na.Get(i)
}

// Elem ...
func (v StrVector) Elem(i int) (interface{}, bool) {
	return v.obs[i], v.na.Get(i)
}

// END TEST ACCESS ------------------------------------------------------------
