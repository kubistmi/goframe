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
			data: []int{1, 2, 0, -20, 6, 9},
			na:   NewNA(),
		}

	} else if kind == "bins" {
		return IntVector{
			data: []int{0, 1, 0, 0, 0, 1, 1, 1, 0},
			na:   NewNA(),
		}
	} else if kind == "uneven" {
		return IntVector{
			data: []int{3, 4, 9, -126, 697, 0, 7, 0, 0, 0},
			na:   NewNA(),
		}
	}

	return IntVector{}
}

func makeStrVec(kind string) StrVector {

	if kind == "uniq" {
		return StrVector{
			data: []string{"bash", "awk", "R", "goframe", "python", "SQL"},
			na:   NewNA(),
		}
	} else if kind == "bins" {
		return StrVector{
			data: []string{"0", "1", "0", "0", "0", "1", "1", "1", "0"},
			na:   NewNA(),
		}
	} else if kind == "uneven" {
		return StrVector{
			data: []string{"a", "z", "b", "Být, či nebýt", "To je oč tu běží", "g", "g", "CDEG", "g", "g"},
			na:   NewNA(),
		}
	}

	return StrVector{}
}
