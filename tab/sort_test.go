package tab

import (
	"testing"

	"github.com/kubistmi/goframe/vec"
)

func TestSort(t *testing.T) {
	df, _ := NewDf(map[string]vec.Vector{
		"age":   vec.NewVec([]int{10, 15, 40, 26, 23, 35, 59, 46}, nil),
		"sex":   vec.NewVec([]string{"m", "f", "f", "f", "m", "f", "m", "m"}, nil),
		"group": vec.NewVec([]int{1, 0, 2, 2, 0, 2, 1, 1}, nil),
	})
	// got :=
	df.Sort([]string{"sex", "group"})

	// fmt.Println(got)
}
