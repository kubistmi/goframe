package tab

import (
	"testing"

	"github.com/kubistmi/goframe/vec"
)

const N = 1000

func BenchmarkSetup(b *testing.B) {
	for n := 0; n < b.N; n++ {
		vals := make([]int, N)
		for i := 0; i < N; i++ {
			vals[i] = i
		}
		NewDf(map[string]vec.Vector{"ints": vec.NewVec(vals, nil)})
	}
}

func BenchmarkMutate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		vals := make([]int, N)
		for i := 0; i < N; i++ {
			vals[i] = i
		}
		df, _ := NewDf(map[string]vec.Vector{"ints": vec.NewVec(vals, nil)})
		df.Mutate(MapF("ints2", func(a int) int { return a * a }, "ints"))
	}
}

func BenchmarkMutateMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		vals := make([]int, N)
		for i := 0; i < N; i++ {
			vals[i] = i
		}
		df, _ := NewDf(map[string]vec.Vector{"ints": vec.NewVec(vals, nil)})
		df.MutateMap(MapF("ints2", func(a map[string]interface{}) int {
			c := a["ints"].(int)
			return c * c
		}, "ints"))
	}

}

// func BenchmarkMutate2(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		vals := make([]int, N)
// 		for i := 0; i < N; i++ {
// 			vals[i] = i
// 		}
// 		df, _ := NewDf(map[string]vec.Vector{"ints": vec.NewVec(vals)})

// 		df.Mutate2(mut2{"ints2", []string{"ints"}, func(args ...interface{}) interface{} {
// 			return args[0].(int) * args[0].(int)
// 		}})
// 	}
// }
