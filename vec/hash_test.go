package vec

import (
	"reflect"
	"testing"
)

func TestIntVector_Hash(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		want Vector
	}

	uniq := makeIntVec("uniq")
	uniqW := makeIntVec("uniq")
	uniqW.hash = hashI{
		lookup: map[int]int{1: 0, 2: 1, 0: 2, -20: 3, 6: 4, 9: 5},
		size:   6,
	}
	bins := makeIntVec("bins")
	binsW := makeIntVec("bins")
	binsW.hash = hashI{
		lookup: map[int]int{0: 0, 1: 1},
		size:   2,
	}
	uneven := makeIntVec("uneven")
	unevenW := makeIntVec("uneven")
	unevenW.hash = hashI{
		lookup: map[int]int{3: 0, 4: 1, 9: 2, -126: 3, 697: 4, 0: 5, 7: 6},
		size:   7,
	}

	tests := []testIntVec{
		testIntVec{"all unique", uniq, uniqW},
		testIntVec{"binary", bins, binsW},
		testIntVec{"binary", uneven, unevenW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Hash(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		want Vector
	}

	uniq := makeStrVec("uniq")
	uniqW := makeStrVec("uniq")
	uniqW.hash = hashS{
		lookup: map[string]int{"bash": 0, "awk": 1, "R": 2, "goframe": 3, "python": 4, "SQL": 5},
		size:   6,
	}
	bins := makeStrVec("bins")
	binsW := makeStrVec("bins")
	binsW.hash = hashS{
		lookup: map[string]int{"0": 0, "1": 1},
		size:   2,
	}
	uneven := makeStrVec("uneven")
	unevenW := makeStrVec("uneven")
	unevenW.hash = hashS{
		lookup: map[string]int{"a": 0, "z": 1, "b": 2, "Být, či nebýt": 3, "To je oč tu běží": 4, "g": 5, "CDEG": 6},
		size:   7,
	}

	tests := []struct {
		name string
		v    StrVector
		want Vector
	}{
		testStrVec{"all unique", uniq, uniqW},
		testStrVec{"binary", bins, binsW},
		testStrVec{"binary", uneven, unevenW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntVector_GetHash(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		args int
		want int
	}

	uniq := makeIntVec("uniq")
	uniq.hash = hashI{
		lookup: map[int]int{1: 0, 2: 1, 0: 2, -20: 3, 6: 4, 9: 5},
		size:   6,
	}
	bins := makeIntVec("bins")
	bins.hash = hashI{
		lookup: map[int]int{0: 0, 1: 1},
		size:   2,
	}

	tests := []testIntVec{
		testIntVec{"unique first", uniq, 1, 0},
		testIntVec{"unique last", uniq, 9, 5},
		testIntVec{"unique middle", uniq, -20, 3},
		testIntVec{"binary zero", bins, 0, 0},
		testIntVec{"binary one", bins, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetHash(tt.args); got != tt.want {
				t.Errorf("IntVector.GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_GetHash(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		args string
		want int
	}

	uniq := makeStrVec("uniq")
	uniq.hash = hashS{
		lookup: map[string]int{"bash": 0, "awk": 1, "R": 2, "goframe": 3, "python": 4, "SQL": 5},
		size:   6,
	}
	bins := makeStrVec("bins")
	bins.hash = hashS{
		lookup: map[string]int{"0": 0, "1": 1},
		size:   2,
	}

	tests := []testStrVec{
		testStrVec{"unique first", uniq, "bash", 0},
		testStrVec{"unique last", uniq, "SQL", 5},
		testStrVec{"unique middle", uniq, "R", 2},
		testStrVec{"binary zero", bins, "0", 0},
		testStrVec{"binary one", bins, "1", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetHash(tt.args); got != tt.want {
				t.Errorf("StrVector.GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntVector_IsHashed(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		want bool
	}

	bins := makeIntVec("bins")
	binsH := makeIntVec("bins")
	binsH.hash = hashI{
		lookup: map[int]int{0: 0, 1: 1},
		size:   2,
	}

	tests := []testIntVec{
		testIntVec{"unhashed", bins, false},
		testIntVec{"hashed", binsH, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsHashed(); got != tt.want {
				t.Errorf("IntVector.IsHashed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_IsHashed(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		want bool
	}

	bins := makeStrVec("bins")
	binsH := makeStrVec("bins")
	binsH.hash = hashS{
		lookup: map[string]int{"0": 0, "1": 1},
		size:   2,
	}

	tests := []testStrVec{
		testStrVec{"unhashed", bins, false},
		testStrVec{"hashed", binsH, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsHashed(); got != tt.want {
				t.Errorf("StrVector.IsHashed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntVector_GetHashVals(t *testing.T) {
	type testIntVec struct {
		name  string
		v     IntVector
		want  []int
		want1 int
	}

	bins := makeIntVec("bins")
	bins.hash = hashI{
		lookup: map[int]int{0: 0, 1: 1},
		size:   2,
	}
	uneven := makeIntVec("uneven")
	uneven.hash = hashI{
		lookup: map[int]int{3: 0, 4: 1, 9: 2, -126: 3, 697: 4, 0: 5, 7: 6},
		size:   7,
	}

	tests := []testIntVec{
		testIntVec{"binary", bins, []int{0, 1, 0, 0, 0, 1, 1, 1, 0}, 2},
		testIntVec{"uneven", uneven, []int{0, 1, 2, 3, 4, 5, 6, 5, 5, 5}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.v.GetHashVals()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.GetHashVals() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IntVector.GetHashVals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStrVector_GetHashVals(t *testing.T) {
	type testStrVec struct {
		name  string
		v     StrVector
		want  []int
		want1 int
	}

	bins := makeStrVec("bins")
	bins.hash = hashS{
		lookup: map[string]int{"0": 0, "1": 1},
		size:   2,
	}
	uneven := makeStrVec("uneven")
	uneven.hash = hashS{
		lookup: map[string]int{"a": 0, "z": 1, "b": 2, "Být, či nebýt": 3, "To je oč tu běží": 4, "g": 5, "CDEG": 6},
		size:   7,
	}

	tests := []testStrVec{
		testStrVec{"binary", bins, []int{0, 1, 0, 0, 0, 1, 1, 1, 0}, 2},
		testStrVec{"uneven", uneven, []int{0, 1, 2, 3, 4, 5, 5, 6, 5, 5}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.v.GetHashVals()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.GetHashVals() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StrVector.GetHashVals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
