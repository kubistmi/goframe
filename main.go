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
	for ix, val := range df.Pulln(0).Loc([]int{1, 2, 5}).GetI().([]int) {
		fmt.Printf("%v = %v\n", ix, val)
	}

	fmt.Println(df.Rows(c(1, 2)))
	fmt.Printf("%v\n", df.Rows(c(20)).err)

	b := df.Mutate(
		mut{"ints", "ints", func(i int) int {
			return i * 3
		}},
		mut{"test", "ints", func(i int) int {
			return i * i
		}})

	fmt.Println(b)

	c := df.Filter(mapf{
		"ints": func(i int) bool {
			return i < 4
		},
		"strs": func(s string) bool {
			return s == "a" || s == "b" || s == "c"
		},
	})
	fmt.Println(c)
}
