package main

import (
	"fmt"
	"log"
)

func inverse(names []string) map[string]int {
	inames := make(map[string]int)
	for ix, val := range names {
		inames[val] = ix
	}
	return inames
}

// Help constructing slices?
func c(p ...int) []int {
	return p
}

type mapf map[string]interface{}

type mut struct {
	new, old string
	fun      interface{}
}

func main() {

	vecI := NewVec([]int{0, 1, 2, 3, 4, 5})
	if vecI.Err() != nil {
		log.Fatal(vecI.Err())
	}
	vecS := NewVec([]string{"a", "b", "c", "d", "e", "f"})
	if vecS.Err() != nil {
		log.Fatal(vecS.Err())
	}
	df, err := NewDf(map[string]Vector{"ints": vecI, "strs": vecS})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df)
	data, _ := df.Pulln(0).Loc([]int{1, 2, 5}).GetI()
	for ix, val := range data.([]int) {
		fmt.Printf("%v = %v\n", ix, val)
		fmt.Printf("%v = %v\n", ix, val)
	}

	fmt.Println(df.Rows(c(1, 2)))
	fmt.Printf("%v\n", df.Rows(c(20)).err)

	a := df.Mutate(
		mut{"ints", "ints", func(i int) int {
			return i * 3
		}},
		mut{"test", "ints", func(i int) int {
			return i * i
		}})
	fmt.Println(a)

	b := df.Filter(mapf{
		"ints": func(i int) bool {
			return i < 4
		},
		"strs": func(s string) bool {
			return s == "a" || s == "b" || s == "c"
		},
	})
	fmt.Println(b)

	c := df.Assign("lints", NewVec([]int{5, 6, 7, 8, 9, 10}))
	fmt.Println(c)

	d := c.Assign("prod", func(a, b Vector) Vector {
		slcA, _ := a.GetI()
		as := slcA.([]int)

		slcB, _ := b.GetI()
		bs := slcB.([]int)
		for ix := range as {
			as[ix] = as[ix] * bs[ix]
		}
		return NewVec(as)
	}(df.Pull("ints"), df.Pull("lints")))

	fmt.Println(d)
}
