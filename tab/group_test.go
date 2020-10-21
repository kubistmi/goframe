package tab

import (
	"reflect"
	"testing"

	"github.com/kubistmi/goframe/vec"
)

func TestGroup(t *testing.T) {
	df, _ := NewDf(map[string]vec.Vector{
		"age":   vec.NewVec([]int{10, 15, 40, 26, 23, 35, 59, 46}),
		"sex":   vec.NewVec([]string{"m", "f", "f", "f", "m", "f", "m", "m"}),
		"group": vec.NewVec([]int{1, 0, 2, 2, 0, 2, 1, 1}),
	})
	got := df.Group([]string{"sex", "group"}).GetGroups()
	want := map[int][]int{0: []int{0, 6, 7}, 2: []int{4}, 3: []int{1}, 5: []int{2, 3, 5}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected: %#v, got: %#v", want, got)
	}

}
