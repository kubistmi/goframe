package utils

func Cumsum() func(int, bool) (int, bool) {
	var na bool
	var sum int
	return func(d int, n bool) (int, bool) {
		if n || na {
			na = true
			return 0, na
		}
		sum += d
		return sum, false
	}
}

func Lag(size int) func(int, bool) (int, bool) {
	na := make(chan bool, size+1)
	buf := make(chan int, size+1)
	c := 0

	return func(d int, n bool) (int, bool) {
		na <- n
		buf <- d
		c++
		if c > size {
			return <-buf, <-na
		}
		return 0, true
	}
}

func Seq(from, to, by int) []int {

	if (by > 0 && from < to) || (by < 0 && from > to) {
		size := (to - from) / by
		out := make([]int, size)
		for ix := range out {
			out[ix] = from + ix*by
		}
		return out
	}

	return []int{}
}

// func Window(size int, out []int) func(int, bool) {
// 	na := NewList(size)
// 	buf := NewList(size)

// 	return func(d int, n bool) {
// 		na.Push(n)
// 		buf.Push(d)
// 		buf.Get(out)
// 	}
// }
