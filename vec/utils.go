package vec

import "fmt"

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
