package vec

import (
	"reflect"
	"testing"
)

func TestIntVector_Mutate(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		args func(v int) int
		want Vector
	}

	uniq := makeIntVec("uniq")
	uniqW := makeIntVec("uniq")
	uniqW.obs = []int{1, 4, 0, 400, 36, 81}

	bins := makeIntVec("bins")
	binsW := makeIntVec("bins")
	binsW.obs = []int{-1, 0, -1, -1, -1, 0, 0, 0, -1}

	tests := []testIntVec{
		testIntVec{"square", uniq, func(v int) int { return v * v }, uniqW},
		testIntVec{"minus one", bins, func(v int) int { return v - 1 }, binsW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mutate(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Mutate(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		args func(v string) string
		want Vector
	}

	uniq := makeStrVec("uniq")
	uniqW := makeStrVec("uniq")
	uniqW.obs = []string{"ba", "aw", "R", "go", "py", "SQ"}

	bins := makeStrVec("bins")
	binsW := makeStrVec("bins")
	binsW.obs = []string{"a0", "a1", "a0", "a0", "a0", "a1", "a1", "a1", "a0"}

	tests := []testStrVec{
		testStrVec{"cut two", uniq, func(v string) string {
			if len(v) > 2 {
				return v[:2]
			} else {
				return v
			}
		}, uniqW},
		testStrVec{"add a", bins, func(v string) string { return "a" + v }, binsW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mutate(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}
