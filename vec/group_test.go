package vec

import (
	"reflect"
	"testing"
)

func TestIntVector_Group(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		want Vector
	}

	uniq := makeIntVec("uniq")
	uniqW := makeIntVec("uniq")
	uniqW.index = map[int][]int{-20: []int{3}, 0: []int{2}, 1: []int{0}, 2: []int{1}, 6: []int{4}, 9: []int{5}}

	bins := makeIntVec("bins")
	binsW := makeIntVec("bins")
	binsW.index = map[int][]int{0: []int{0, 2, 3, 4, 8}, 1: []int{1, 5, 6, 7}}

	uneven := makeIntVec("uneven")
	unevenW := makeIntVec("uneven")
	unevenW.index = map[int][]int{3: []int{0}, 4: []int{1}, 9: []int{2}, -126: []int{3}, 697: []int{4}, 0: []int{5, 7, 8, 9}, 7: []int{6}}

	tests := []testIntVec{
		testIntVec{"all unique", uniq, uniqW},
		testIntVec{"binary", bins, binsW},
		testIntVec{"binary", uneven, unevenW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Group(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Group() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Group(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		want Vector
	}

	uniq := makeStrVec("uniq")
	uniqW := makeStrVec("uniq")
	uniqW.index = map[string][]int{"bash": []int{0}, "awk": []int{1}, "R": []int{2}, "goframe": []int{3}, "python": []int{4}, "SQL": []int{5}}

	bins := makeStrVec("bins")
	binsW := makeStrVec("bins")
	binsW.index = map[string][]int{"0": []int{0, 2, 3, 4, 8}, "1": []int{1, 5, 6, 7}}

	uneven := makeStrVec("uneven")
	unevenW := makeStrVec("uneven")
	unevenW.index = map[string][]int{"a": []int{0}, "z": []int{1}, "b": []int{2}, "Být, či nebýt": []int{3}, "To je oč tu běží": []int{4}, "g": []int{5, 6, 8, 9}, "CDEG": []int{7}}

	tests := []testStrVec{
		testStrVec{"all unique", uniq, uniqW},
		testStrVec{"binary", bins, binsW},
		testStrVec{"binary", uneven, unevenW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Group(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Group() = %v, want %v", got, tt.want)
			}
		})
	}
}
