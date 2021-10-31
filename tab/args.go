package tab

import (
	"fmt"
	"reflect"
)

type Arg struct {
	left  []string
	right []string
	kind  string
	err   error
}

type Option func(a *Arg)

func Options(opts ...Option) Option {
	return func(a *Arg) {
		for _, opt := range opts {
			opt(a)
		}
	}
}

func OnMap(m map[string]string) func(*Arg) {
	left := make([]string, 0, len(m))
	right := make([]string, 0, len(m))
	for x, y := range m {
		left = append(left, x)
		right = append(right, y)
	}

	return func(a *Arg) {
		a.left = left
		a.right = right
	}
}

func On(kind string, cols []string) func(*Arg) {
	if kind == "" {
		return func(a *Arg) {
			a.left = cols
			a.right = cols
		}
	} else if kind == "left" {
		return func(a *Arg) {
			a.left = cols
		}
	} else if kind == "right" {
		return func(a *Arg) {
			a.right = cols
		}
	} else {
		return func(a *Arg) {
			a.err = fmt.Errorf("Wrong kind specification, expected one of: 'right', 'left', '', got: %v", kind)
		}
	}
}

// Reflection and other tricks!
type Args struct {
	Number int
	By     string
}

func By(c string) func(interface{}) {
	return func(i interface{}) {
		reflect.ValueOf(i).Elem().FieldByName("By").SetString(c)
	}
}

// foo := Args{123, "Hello"}
// change := By("test")
// change(&foo)
// reflect.ValueOf(&foo).Elem().Field(0).SetInt(321)
// fmt.Println(foo)
