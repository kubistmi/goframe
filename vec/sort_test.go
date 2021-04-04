package vec

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOrder(t *testing.T) {
	vec := NewVec([]int{1, 6, 22, 4, 9, 7, 8, 9}, NewNA())
	ix := vec.(IntVector).Order()
	fmt.Println(vec)
	fmt.Println(ix)
	fmt.Println(vec.Loc(ix))
}

func TestIntVector_Sort(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		want Vector
	}

	ordr := IntVector{
		data: []int{0, 1, 2, 3, 5, 4},
	}
	ordrS := IntVector{
		data: []int{0, 1, 2, 3, 4, 5},
	}

	uniq := makeIntVec("uniq")
	uniqS := makeIntVec("uniq")
	uniqS.data = []int{-20, 0, 1, 2, 6, 9}

	bins := makeIntVec("bins")
	binsS := makeIntVec("bins")
	binsS.data = []int{0, 0, 0, 0, 0, 1, 1, 1, 1}

	tests := []testIntVec{
		testIntVec{"ordr", ordr, ordrS},
		testIntVec{"unique", uniq, uniqS},
		testIntVec{"binary", bins, binsS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Sort(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		want Vector
	}

	ordr := StrVector{
		data: []string{"a", "b", "c", "d"},
	}

	rev := StrVector{
		data: []string{"z", "y", "x", "w", "v", "a"},
	}
	revS := StrVector{
		data: []string{"a", "v", "w", "x", "y", "z"},
	}

	uniq := makeStrVec("uniq")
	uniqS := makeStrVec("uniq")
	uniqS.data = []string{"R", "SQL", "awk", "bash", "goframe", "python"}

	bins := makeStrVec("bins")
	binsS := makeStrVec("bins")
	binsS.data = []string{"0", "0", "0", "0", "0", "1", "1", "1", "1"}

	tests := []testStrVec{
		testStrVec{"ordered", ordr, ordr},
		testStrVec{"reversed", rev, revS},
		testStrVec{"unique", uniq, uniqS},
		testStrVec{"binary", bins, binsS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntVector_Order(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		want []int
	}

	uniq := makeIntVec("uniq")

	uneven := makeIntVec("uneven")

	tests := []testIntVec{
		testIntVec{"unique", uniq, []int{3, 2, 0, 1, 4, 5}},
		testIntVec{"uneven", uneven, []int{3, 5, 7, 8, 9, 0, 1, 6, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Order(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Order(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		want []int
	}

	uniq := makeStrVec("uniq")

	uneven := makeStrVec("uneven")

	tests := []testStrVec{
		testStrVec{"unique", uniq, []int{2, 5, 1, 0, 3, 4}},
		testStrVec{"uneven", uneven, []int{3, 7, 4, 0, 2, 5, 6, 8, 9, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Order(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Order() = %v, want %v", got, tt.want)
			}
		})
	}
}
