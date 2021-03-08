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
	return fmt.Errorf("%w parameter f, expected: `func(T) T` OR `func(T) bool` for T = `string` / `int`, got: `%T`", ErrParamType, f)
}

//! I am so ready for the generics!
// func SkipNA[I, O any](f func(I) O) func(I, bool) (O, bool) {

// 	switch (interface{})(f).(type) {
// 	case func(I) I:
// 		fmt.Println("mutate")

// 	case func(I) bool:
// 		fmt.Println("check")
// 	}

// 	return func(d I, na bool) (O, bool) {
// 		if na {
// 			var e O
// 			return e, true
// 		}
// 		return f(d), false
// 	}
// }
