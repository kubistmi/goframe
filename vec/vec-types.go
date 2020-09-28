package vec

// IntVector ... ---------------------------------------------------------------
type IntVector struct {
	obs   []int
	na    Set
	index struct {
		lookup map[int]int
		size   int
	}
	size int
	err  error
}

// Size ...
func (v IntVector) Size() int {
	return v.size
}

// StrVector ... ---------------------------------------------------------------
type StrVector struct {
	obs   []string
	na    Set
	index struct {
		lookup map[string]int
		size   int
	}
	size    int
	inverse map[string][]int //! inverse index
	err     error
}

// Size ...
func (v StrVector) Size() int {
	return v.size
}
