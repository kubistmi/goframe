package vec

// Elem ... TEST ACCESS -------------------------------------------------------
func (v IntVector) Elem(i int) (interface{}, bool) {
	return v.obs[i], v.na[i]
}

// Elem ...
func (v StrVector) Elem(i int) (interface{}, bool) {
	return v.obs[i], v.na[i]
}

// END TEST ACCESS ------------------------------------------------------------
