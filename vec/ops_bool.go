package vec

// Datatype describes the type of the Vector
type comparator int

const (
	gt comparator = iota
	ge comparator = iota
	lt comparator = iota
	le comparator = iota
	eq comparator = iota
)

func (v IntVector) Gt(a int) BoolVector {
	return BoolVector{}
}

func (v IntVector) Ge(a int) BoolVector {
	return BoolVector{}
}

func (v IntVector) Lt(a int) BoolVector {
	return BoolVector{}
}

func (v IntVector) Le(a int) BoolVector {
	return BoolVector{}
}

func (v IntVector) Eq(a int) BoolVector {
	return BoolVector{}
}

func (v StrVector) Gt(a string) BoolVector {
	return BoolVector{}
}

func (v StrVector) Ge(a string) BoolVector {
	return BoolVector{}
}

func (v StrVector) Lt(a string) BoolVector {
	return BoolVector{}
}

func (v StrVector) Le(a string) BoolVector {
	return BoolVector{}
}

func (v StrVector) Eq(a string) BoolVector {
	return BoolVector{}
}

// func runBoolOpStr()
// func runBoolOpInt()

func compareInt(a, b string, kind comparator) bool {
	switch kind {
	case gt:
		return a > b
	case ge:
		return a >= b
	case lt:
		return a < b
	case le:
		return a <= b
	case eq:
		return a == b
	}
	return false
}

func compareStr(a, b string, kind comparator) bool {
	switch kind {
	case gt:
		return a > b
	case ge:
		return a >= b
	case lt:
		return a < b
	case le:
		return a <= b
	case eq:
		return a == b
	}
	return false
}
