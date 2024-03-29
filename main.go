package main

import (
	"log"

	"github.com/kubistmi/goframe/tab"
	"github.com/kubistmi/goframe/vec"
)

func main() {

	vecI := vec.NewVec([]int{0, 2, 2, 5, 4, 5}, vec.NewNA(0))
	if vecI.Err() != nil {
		log.Fatal(vecI.Err())
	}
	vecS := vec.NewVec([]string{"a", "c", "b", "z", "b", "a"}, vec.NewNA(0))
	if vecS.Err() != nil {
		log.Fatal(vecS.Err())
	}
	df, err := tab.NewDf(map[string]vec.Vector{"ints": vecI, "strs": vecS})
	if err != nil {
		log.Fatal(err)
	}
	df = df.Assign("vals", vec.NewIntVec([]int{1, 2, 3, 4, 5, 6}, vec.NewNA(0)))

	//fmt.Println(df)
	//fmt.Println(df.Spread([]string{"strs"}, "ints", "vals"))
	df.PrintTable()

	// fmt.Println(df.Sort([]string{"ints", "strs"}))

	// fmt.Println(vecI.Hash().GetHashVals())

	// ixed := df.Group([]string{"strs"})
	// fmt.Println(ixed)

	// eg := df.Mutate2(mut2{"res", []string{"ints", "strs"}, func(args ...interface{}) interface{} {
	// 	i := args[0].(int)
	// 	return i * i
	// }})

	// fmt.Println(eg)
	// data, _ := df.Pulln(0).Loc([]int{1, 2, 5}).GetI()
	// for ix, val := range data.([]int) {
	// 	fmt.Printf("%v = %v\n", ix, val)
	// 	fmt.Printf("%v = %v\n", ix, val)
	// }

	// fmt.Println(df.Rows(c(1, 2)))
	// fmt.Printf("%v\n", df.Rows(c(20)).err)

	// a := df.Mutate(
	// 	mut{"ints", "ints", func(i int) int {
	// 		return i * 3
	// 	}},
	// 	mut{"test", "ints", func(i int) int {
	// 		return i * i
	// 	}})
	// fmt.Println(a)

	// b := df.Filter(mapf{
	// 	"ints": func(i int) bool {
	// 		return i < 4
	// 	},
	// 	"strs": func(s string) bool {
	// 		return s == "a" || s == "b" || s == "c"
	// 	},
	// })
	// fmt.Println(b)

	// c := df.Assign("lints", NewVec([]int{5, 6, 7, 8, 9, 10}))
	// fmt.Println(c)

	// d := c.Assign("prod", func(a, b Vector) Vector {
	// 	slcA, _ := a.GetI()
	// 	as := slcA.([]int)

	// 	slcB, _ := b.GetI()
	// 	bs := slcB.([]int)
	// 	for ix := range as {
	// 		as[ix] = as[ix] * bs[ix]
	// 	}
	// 	return NewVec(as)
	// }(df.Pull("ints"), df.Pull("lints")))

	// fmt.Println(d)

	// df, _ := NewDf(map[string]Vector{
	// 	"age":   NewVec([]int{10, 15, 40, 26, 23, 35, 59, 46}),
	// 	"sex":   NewVec([]string{"m", "f", "f", "f", "m", "f", "m", "m"}),
	// 	"group": NewVec([]int{1, 0, 2, 2, 0, 2, 1, 1}),
	// })

	// fmt.Println(df.Group([]string{"sex", "group"}).GetIndex())

	// N := 10000000
	// vals := make([]int, N)
	// for i := 0; i < N; i++ {
	// 	vals[i] = i
	// }
	// df, _ := tab.NewDf(map[string]vec.Vector{"ints": vec.NewVec(vals, nil)})

	// before := time.Now()
	// df.Mutate(tab.MapF("ints2", func(a int, c bool) (int, bool) { return a * a, true }, "ints"))
	// fmt.Printf("Mutate:  %v\n", time.Since(before).Milliseconds())

	// before = time.Now()
	// df.MutateMap(tab.MapF("ints2", func(a map[string]interface{}) int {
	// 	c := a["ints"].(int)
	// 	return c * c
	// }, "ints"))
	// fmt.Printf("MutateMap:  %v\n", time.Since(before).Milliseconds())

	// before = time.Now()
	// df.Mutate2(mut2{"ints2", []string{"ints"}, func(args ...interface{}) interface{} {
	// 	a := args[0].(int)
	// 	return a * a
	// }})
	// fmt.Printf("Mutate2: %v\n", time.Now().Sub(before).Milliseconds())

	// var a interface{}
	// a = func([]int, []string) string {
	// 	return ""
	// }

	// ab(a)
}
