package vec

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kubistmi/goframe/utils"
)

func TestIntVector_Mutate(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		args interface{}
		want Vector
	}

	uniq := makeIntVec("uniq")
	uniqW := makeIntVec("uniq")
	uniqW.data = []int{1, 4, 0, 400, 36, 81}

	bins := makeIntVec("bins")
	binsW := makeIntVec("bins")
	binsW.data = []int{-1, 0, -1, -1, -1, 0, 0, 0, -1}

	errW := IntVector{err: fmt.Errorf("wrong function, expected: `func(int) int`, got: `%w`", fmt.Errorf("undefined function specification"))}

	tests := []testIntVec{
		testIntVec{"square", uniq, utils.SkipNA(func(v int) int { return v * v }), uniqW},
		testIntVec{"minus one", bins, utils.SkipNA(func(v int) int { return v - 1 }), binsW},
		testIntVec{"error", bins, utils.SkipNA(func(v string) int { return 5 }), errW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mutate(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Mutate() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestStrVector_Mutate(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		args interface{}
		want Vector
	}

	uniq := makeStrVec("uniq")
	uniqW := makeStrVec("uniq")
	uniqW.data = []string{"ba", "aw", "R", "go", "py", "SQ"}

	bins := makeStrVec("bins")
	binsW := makeStrVec("bins")
	binsW.data = []string{"a0", "a1", "a0", "a0", "a0", "a1", "a1", "a1", "a0"}

	errW := StrVector{err: fmt.Errorf("wrong function, expected: `func(string) string`, got: `%w`", fmt.Errorf("undefined function specification"))}

	tests := []testStrVec{
		testStrVec{"cut two", uniq, utils.SkipNA(func(v string) string {
			if len(v) > 2 {
				return v[:2]
			}
			return v
		}), uniqW},
		testStrVec{"add a", bins, utils.SkipNA(func(v string) string { return "a" + v }), binsW},
		testStrVec{"error", bins, utils.SkipNA(func(v string) int { return 2 }), errW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mutate(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Mutate() = %s, want %s", got, tt.want)
			}
		})
	}
}
