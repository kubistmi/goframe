package vec

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kubistmi/goframe/utils"
)

func TestIntVector_Loc(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		args []int
		want Vector
	}

	uniq := makeIntVec("uniq")
	first := makeIntVec("uniq")
	first.obs = first.obs[:3]
	first.size = 3

	last := makeIntVec("uniq")
	last.obs = last.obs[4:]
	last.size = 2

	middle := makeIntVec("uniq")
	middle.obs = middle.obs[2:4]
	middle.size = 2

	errWs := NewErrVec(fmt.Errorf("wrong position, maximum allowed: 5, got -1"))
	errWl := NewErrVec(fmt.Errorf("wrong position, maximum allowed: 5, got 10"))

	tests := []testIntVec{
		testIntVec{"first three", uniq, []int{0, 1, 2}, first},
		testIntVec{"last two", uniq, []int{4, 5}, last},
		testIntVec{"middle two", uniq, []int{2, 3}, middle},
		testIntVec{"error smaller", uniq, []int{-1}, errWs},
		testIntVec{"error larger", uniq, []int{10}, errWl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Loc(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Loc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Loc(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		args []int
		want Vector
	}

	uniq := makeStrVec("uniq")
	first := makeStrVec("uniq")
	first.obs = first.obs[:3]
	first.size = 3

	last := makeStrVec("uniq")
	last.obs = last.obs[4:]
	last.size = 2

	middle := makeStrVec("uniq")
	middle.obs = middle.obs[2:4]
	middle.size = 2

	errWs := NewErrVec(fmt.Errorf("wrong position, maximum allowed: 5, got -1"))
	errWl := NewErrVec(fmt.Errorf("wrong position, maximum allowed: 5, got 10"))

	tests := []testStrVec{
		testStrVec{"first three", uniq, []int{0, 1, 2}, first},
		testStrVec{"last two", uniq, []int{4, 5}, last},
		testStrVec{"middle two", uniq, []int{2, 3}, middle},
		testStrVec{"error smaller", uniq, []int{-1}, errWs},
		testStrVec{"error larger", uniq, []int{10}, errWl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Loc(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Loc() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestIntVector_Check(t *testing.T) {
	type testIntVec struct {
		name    string
		v       IntVector
		args    interface{}
		want    []bool
		wantErr error
	}
	T := true
	F := false

	uniq := makeIntVec("uniq")
	uneven := makeIntVec("uneven")
	bins := makeIntVec("bins")

	tests := []testIntVec{
		testIntVec{"single", uneven, utils.SkipNA(func(v int) bool { return v == 3 }), []bool{T, F, F, F, F, F, F, F, F, F}, nil},
		testIntVec{"multiple", bins, utils.SkipNA(func(v int) bool { return v == 0 }), []bool{T, F, T, T, T, F, F, F, T}, nil},
		testIntVec{"none", uniq, utils.SkipNA(func(v int) bool { return v > 30 }), []bool{F, F, F, F, F, F}, nil},
		testIntVec{"error", uniq, utils.SkipNA(func(v string) bool { return v > "A" }), nil, fmt.Errorf("wrong function, expected: `func(int) bool`, got: `func(string, bool) bool`")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Check(tt.args)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("IntVector.Check() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Check(t *testing.T) {
	type testStrVec struct {
		name    string
		v       StrVector
		args    interface{}
		want    []bool
		wantErr error
	}
	T := true
	F := false

	uniq := makeStrVec("uniq")
	bins := makeStrVec("bins")

	tests := []testStrVec{
		testStrVec{"single", uniq, utils.SkipNA(func(v string) bool { return v == "R" }), []bool{F, F, T, F, F, F}, nil},
		testStrVec{"multiple", bins, utils.SkipNA(func(v string) bool { return v == "0" }), []bool{T, F, T, T, T, F, F, F, T}, nil},
		testStrVec{"none", bins, utils.SkipNA(func(v string) bool { return v > "z" }), []bool{F, F, F, F, F, F, F, F, F}, nil},
		testStrVec{"error", uniq, utils.SkipNA(func(v string) int { return 0 }), nil, fmt.Errorf("wrong function, expected: `func(string) bool`, got: `%w`", fmt.Errorf("undefined function specification"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Check(tt.args)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StrVector.Check() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntVector_Filter(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		args interface{}
		want Vector
	}

	uniq := makeIntVec("uniq")
	betw := makeIntVec("uniq")
	betw.obs = betw.obs[:3]
	betw.size = 3

	last := makeIntVec("uniq")
	last.obs = last.obs[4:]
	last.size = 2

	bins := makeIntVec("bins")
	binsW := makeIntVec("bins").Loc([]int{0, 2, 3, 4, 8})

	errW := StrVector{err: fmt.Errorf(
		"error in Check: %w",
		fmt.Errorf("wrong function, expected: `func(int) bool`, got: `%w`",
			fmt.Errorf("undefined function specification")),
	)}

	tests := []testIntVec{
		testIntVec{"between", uniq, utils.SkipNA(func(v int) bool { return v > -1 && v < 3 }), betw},
		testIntVec{"higher", uniq, utils.SkipNA(func(v int) bool { return v >= 6 }), last},
		testIntVec{"equals", bins, utils.SkipNA(func(v int) bool { return v == 0 }), binsW},
		testIntVec{"error", bins, utils.SkipNA(func(a string) int { return 2 }), errW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Filter(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Filter() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestStrVector_Filter(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		args interface{}
		want Vector
	}

	uniq := makeStrVec("uniq")
	smaller := makeStrVec("uniq")
	smaller.obs = smaller.obs[:2]
	smaller.size = 2

	last := makeStrVec("uniq")
	last.obs = last.obs[4:5]
	last.size = 1

	bins := makeStrVec("bins")
	binsW := makeStrVec("bins").Loc([]int{1, 5, 6, 7})

	uneven := makeStrVec("uneven")
	isin := makeStrVec("uneven")
	isin.obs = isin.obs[:3]
	isin.size = 3

	errW := StrVector{err: fmt.Errorf(
		"error in Check: %w",
		fmt.Errorf("wrong function, expected: `func(string) bool`, got: `%w`",
			fmt.Errorf("undefined function specification")),
	)}

	tests := []testStrVec{
		testStrVec{"between", uniq, utils.SkipNA(func(v string) bool { return v > "a" && v < "bb" }), smaller},
		testStrVec{"larger", uniq, utils.SkipNA(func(v string) bool { return v > "p" }), last},
		testStrVec{"equals", bins, utils.SkipNA(func(v string) bool { return v == "1" }), binsW},
		testStrVec{"isin", uneven, utils.SkipNA(Isin([]string{"a", "b", "z"})), isin},
		testStrVec{"error", bins, utils.SkipNA(func(a int) string { return "nope" }), errW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Filter(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Filter() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestIntVector_Mask(t *testing.T) {
	type testIntVec struct {
		name string
		v    IntVector
		args []bool
		want Vector
	}
	T := true
	F := false

	uniq := makeIntVec("uniq")
	first := makeIntVec("uniq")
	first.obs = first.obs[:3]
	first.size = 3

	last := makeIntVec("uniq")
	last.obs = last.obs[4:]
	last.size = 2

	middle := makeIntVec("uniq")
	middle.obs = middle.obs[2:4]
	middle.size = 2

	errWs := NewErrVec(fmt.Errorf("size of boolean slice does not match the size of Vector, expected: 6, got: 2"))
	errWl := NewErrVec(fmt.Errorf("size of boolean slice does not match the size of Vector, expected: 6, got: 8"))

	tests := []testIntVec{
		testIntVec{"first three", uniq, []bool{T, T, T, F, F, F}, first},
		testIntVec{"last two", uniq, []bool{F, F, F, F, T, T}, last},
		testIntVec{"middle two", uniq, []bool{F, F, T, T, F, F}, middle},
		testIntVec{"error smaller", uniq, []bool{F, F}, errWs},
		testIntVec{"error larger", uniq, []bool{F, F, T, T, F, F, T, T}, errWl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mask(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntVector.Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVector_Mask(t *testing.T) {
	type testStrVec struct {
		name string
		v    StrVector
		args []bool
		want Vector
	}
	T := true
	F := false

	uniq := makeStrVec("uniq")
	first := makeStrVec("uniq")
	first.obs = first.obs[:3]
	first.size = 3

	last := makeStrVec("uniq")
	last.obs = last.obs[4:]
	last.size = 2

	middle := makeStrVec("uniq")
	middle.obs = middle.obs[2:4]
	middle.size = 2

	errWs := NewErrVec(fmt.Errorf("size of boolean slice does not match the size of Vector, expected: 6, got: 2"))
	errWl := NewErrVec(fmt.Errorf("size of boolean slice does not match the size of Vector, expected: 6, got: 8"))

	tests := []testStrVec{
		testStrVec{"first three", uniq, []bool{T, T, T, F, F, F}, first},
		testStrVec{"last two", uniq, []bool{F, F, F, F, T, T}, last},
		testStrVec{"middle two", uniq, []bool{F, F, T, T, F, F}, middle},
		testStrVec{"error smaller", uniq, []bool{F, F}, errWs},
		testStrVec{"error larger", uniq, []bool{F, F, T, T, F, F, T, T}, errWl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mask(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrVector.Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
