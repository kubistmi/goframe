package tab

import (
	"fmt"

	"github.com/kubistmi/goframe/vec"
)

type mut struct {
	new, old string
	fun      interface{}
}

// Mutate ///
func (df Table) Mutate(mf ...mut) Table {

	for _, val := range mf {
		switch v := df.data[val.old].(type) {
		case vec.IntVector:
			switch f := val.fun.(type) {
			case func(int) int:
				if _, ok := df.data[val.new]; !ok {
					df.names = append(df.names, val.new)
				}
				df.data[val.new] = v.Mutate(f)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		case vec.StrVector:
			switch f := val.fun.(type) {
			case func(string) string:
				if _, ok := df.data[val.new]; !ok {
					df.names = append(df.names, val.new)
				}
				df.data[val.new] = v.Mutate(f)
			default:
				return Table{
					err: fmt.Errorf("wrong function definition, expected func(int) int, got %T", f),
				}
			}
		}
	}
	return df
}

// Pipe ...
func (df Table) Pipe(f func(Table) Table) Table {
	return f(df)
}
