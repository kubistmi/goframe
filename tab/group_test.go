package tab

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kubistmi/goframe/vec"
)

func meanAge(v vec.Vector) int {
	vals, _ := v.(vec.IntVector).Get()
	sum := 0
	for _, x := range vals {
		sum = sum + x
	}
	return sum / len(vals)
}

func TestGroup(t *testing.T) {
	df, _ := NewDf(map[string]vec.Vector{
		"age":   vec.NewVec([]int{10, 15, 40, 26, 23, 35, 59, 46}),
		"sex":   vec.NewVec([]string{"m", "f", "f", "f", "m", "f", "m", "m"}),
		"group": vec.NewVec([]int{1, 0, 2, 2, 0, 2, 1, 1}),
	})
	grDf := df.Group([]string{"sex", "group"})
	got := grDf.GetGroups()
	want := map[int][]int{0: []int{0, 6, 7}, 2: []int{4}, 3: []int{1}, 5: []int{2, 3, 5}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected: %#v, got: %#v", want, got)
	}

	fmt.Println(grDf.Agg(map[string]interface{}{"age": meanAge}))
}
