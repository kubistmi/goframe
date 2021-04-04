package vec

type hashI struct {
	lookup map[int]int
	size   int
}

type hashS struct {
	lookup map[string]int
	size   int
}

func makeIntVec(kind string) IntVector {

	if kind == "uniq" {
		return NewIntVec([]int{1, 2, 0, -20, 6, 9}, NewNA())

	} else if kind == "bins" {
		return NewIntVec([]int{0, 1, 0, 0, 0, 1, 1, 1, 0}, NewNA())

	} else if kind == "uneven" {
		return NewIntVec([]int{3, 4, 9, -126, 697, 0, 7, 0, 0, 0}, NewNA())
	}
	return NewIntVec([]int{}, nil)
}

func makeStrVec(kind string) StrVector {

	if kind == "uniq" {
		return NewStrVec([]string{"bash", "awk", "R", "goframe", "python", "SQL"}, NewNA())

	} else if kind == "bins" {
		return NewStrVec([]string{"0", "1", "0", "0", "0", "1", "1", "1", "0"}, NewNA())

	} else if kind == "uneven" {
		return NewStrVec([]string{"a", "z", "b", "Být, či nebýt", "To je oč tu běží", "g", "g", "CDEG", "g", "g"}, NewNA())

	}

	return NewStrVec([]string{}, nil)
}
