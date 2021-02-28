package vec

import "fmt"

// Isin ...
func Isin(w interface{}) interface{} {

	var e struct{}

	switch t := w.(type) {
	case []int:
		l := make(map[int]struct{}, len(t))
		for _, val := range t {
			l[val] = e
		}
		return func(v int) bool {
			_, ok := l[v]
			return ok
		}

	case []string:
		l := make(map[string]struct{}, len(t))
		for _, val := range t {
			l[val] = e
		}
		return func(v string) bool {
			_, ok := l[v]
			return ok
		}
	}

	return fmt.Errorf("wrong parameter type, expected: []int/[]string, got: %T", w)

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
