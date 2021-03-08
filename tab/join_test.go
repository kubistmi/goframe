package tab

import (
	"reflect"
	"testing"

	"github.com/kubistmi/goframe/vec"
)

func TestTable_LeftJoin(t *testing.T) {

	type testTable struct {
		name  string
		left  Table
		right Table
		on    []string
		want  Table
	}

	x, _ := NewDf(map[string]vec.Vector{
		"id":  vec.NewVec([]int{0, 6, 7, 1, 2}),
		"val": vec.NewVec([]string{"a", "b", "c", "d", "e"}),
	})

	y, _ := NewDf(map[string]vec.Vector{
		"id":   vec.NewVec([]int{1, 2, 0, 7, 6}),
		"val2": vec.NewVec([]string{"a", "b", "c", "d", "e"}),
	})

	tests := []testTable{
		testTable{"common", x, y, []string{"id"}, Table{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.left.LeftJoin(tt.right, tt.on); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.LeftJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
