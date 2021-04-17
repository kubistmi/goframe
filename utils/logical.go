package utils

import "fmt"

func Any(v interface{}) (bool, error) {
	var out bool = false
	switch b := v.(type) {
	case []bool:
		for _, val := range b {
			out = out || val
		}
	case map[string]bool:
		for _, val := range b {
			out = out || val
		}
	default:
		return false, fmt.Errorf("wrong type")
	}

	return out, nil
}

func All(v interface{}) (bool, error) {
	var out bool = true
	switch b := v.(type) {
	case []bool:
		for _, val := range b {
			out = out && val
		}
	case map[string]bool:
		out = false
		for _, val := range b {
			out = out && val
		}
	default:
		return false, fmt.Errorf("wrong type")
	}

	return out, nil
}

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

	return fmt.Errorf("%w parameter w, expected: `[]int` / `[]string`, got: %T", ErrParamType, w)

}
