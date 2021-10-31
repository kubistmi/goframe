package tab

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kubistmi/goframe/vec"
)

func meanAge(v vec.Vector) int {
	vals, _, err := v.(vec.IntVector).Get()
	if err != nil {
		return 0
	}
	sum := 0
	for _, x := range vals {
		sum = sum + x
	}
	return sum / len(vals)
}

func TestGroup(t *testing.T) {
	df, _ := NewDf(map[string]vec.Vector{
		"age":   vec.NewVec([]int{10, 15, 40, 26, 23, 35, 59, 46}, nil),
		"sex":   vec.NewVec([]string{"m", "f", "f", "f", "m", "f", "m", "m"}, nil),
		"group": vec.NewVec([]int{1, 0, 2, 2, 0, 2, 1, 1}, nil),
	})
	grDf := df.Group([]string{"sex", "group"})
	got := grDf.GroupGet()
	want := map[int][]int{4: {0, 6, 7}, 7: {4}, 8: {1}, 11: {2, 3, 5}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected: %#v, got: %#v", want, got)
	}

	fmt.Println(grDf.Agg(MapF("age", meanAge)))
}
