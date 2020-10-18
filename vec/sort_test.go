package vec

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	vec := NewVec([]int{1, 6, 22, 4, 9, 7, 8, 9})
	sorted := vec.(IntVector).Sort()

	fmt.Println(sorted)
	fmt.Println(vec)
}

func TestOrder(t *testing.T) {
	vec := NewVec([]int{1, 6, 22, 4, 9, 7, 8, 9})
	ix := vec.(IntVector).Order()
	fmt.Println(vec)
	fmt.Println(ix)
	fmt.Println(vec.Loc(ix))
}
