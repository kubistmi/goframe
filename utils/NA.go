package utils

import "fmt"

// SkipNA ...
func SkipNA(f interface{}) interface{} {
	switch fun := f.(type) {

	//vec.Check
	case func(int) bool:
		return func(v int, na bool) bool {
			if na {
				return false
			}
			return fun(v)
		}
	//vec.Check
	case func(string) bool:
		return func(v string, na bool) bool {
			if na {
				return false
			}
			return fun(v)
		}

	//vec.Mutate
	case func(int) int:
		return func(v int, na bool) (int, bool) {
			if na {
				return 0, true
			}
			return fun(v), false
		}
	//vec.Mutate
	case func(string) string:
		return func(v string, na bool) (string, bool) {
			if na {
				return "", true
			}
			return fun(v), false
		}
	}
	return fmt.Errorf("undefined function specification")
}
