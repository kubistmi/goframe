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
		return IntVector{
			obs:  []int{1, 2, 0, -20, 6, 9},
			size: 6,
			na:   make(Set),
		}

	} else if kind == "bins" {
		return IntVector{
			obs:  []int{0, 1, 0, 0, 0, 1, 1, 1, 0},
			size: 9,
			na:   make(Set),
		}
	} else if kind == "uneven" {
		return IntVector{
			obs:  []int{3, 4, 9, -126, 697, 0, 7, 0, 0, 0},
			size: 10,
			na:   make(Set),
		}
	}

	return IntVector{}
}

func makeStrVec(kind string) StrVector {

	if kind == "uniq" {
		return StrVector{
			obs:  []string{"bash", "awk", "R", "goframe", "python", "SQL"},
			size: 6,
			na:   make(Set),
		}
	} else if kind == "bins" {
		return StrVector{
			obs:  []string{"0", "1", "0", "0", "0", "1", "1", "1", "0"},
			size: 9,
			na:   make(Set),
		}
	} else if kind == "uneven" {
		return StrVector{
			obs:  []string{"a", "z", "b", "Být, či nebýt", "To je oč tu běží", "g", "g", "CDEG", "g", "g"},
			size: 10,
			na:   make(Set),
		}
	}

	return StrVector{}
}
