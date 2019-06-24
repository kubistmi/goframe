package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Row struct {
	Names map[string]int
	Data  []float64
}

type Df struct {
	//Shape []int
	Rows []Row
}

type Subquery struct {
	LS, RS, Operator string
}

func (r *Row) L(col string) float64 {
	ix := r.Names[col]
	return r.Data[ix]
}

func ParseQuery(query string) []Subquery {
	query = strings.Replace(query, " ", "", -1)
	parts := strings.Split(query, "&")

	rOp := regexp.MustCompile("(<={0,1}|>={0,1}|=)")

	subqueries := make([]Subquery, len(parts))

	for _, val := range parts {
		var sub Subquery
		sub.Operator = rOp.FindStringSubmatch(val)[0]
		if len(sub.Operator) > 1 {
			log.Fatalf("Too many operators, expected one, got %s.\n", sub.Operator)
		}
		sides := strings.Split(val, sub.Operator)
		if len(sides) > 2 {
			log.Fatalf("Too many parts of the query. The expectation are [left right] (2 sides around the operator), got %s.\n", sides)
		}

		sub.LS = sides[0]
		sub.RS = sides[1]
		subqueries = append(subqueries, sub)
	}
	return subqueries
}

func GetOperator(s Subquery, eq bool) func(r *Row, s Subquery, eq bool) bool {
	operator := s.Operator
	operator = strings.Replace(operator, "=", "", 1)

	switch operator {
	case "<":
		return func(r *Row, s Subquery, eq bool) bool {
			col := s.LS
			val, err := strconv.ParseFloat(s.RS, 64)
			if err != nil {
				log.Fatal(err)
			}
			if eq {
				return r.L(col) <= val
			}
			return r.L(col) < val
		}
	case ">":
		return func(r *Row, s Subquery, eq bool) bool {
			col := s.LS
			val, err := strconv.ParseFloat(s.RS, 64)
			if err != nil {
				log.Fatal(err)
			}
			if eq {
				return r.L(col) >= val
			}
			return r.L(col) > val
		}
	default:
		return func(r *Row, s Subquery, eq bool) bool {
			col := s.LS
			val, err := strconv.ParseFloat(s.RS, 64)
			if err != nil {
				log.Fatal(err)
			}
			return r.L(col) == val
		}
	}
}

func remove(s []float64, i int) []float64 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (d Df) F(query string) Df {
	sub := ParseQuery(query)
	sub1 := sub[1] //!SHIT

	var operations []func(r *Row, s Subquery, eq bool) bool

	//for ix := range sub {
	//	equal := strings.Contains(sub[ix].Operator, "=")
	//	operations = append(operations, GetOperator(sub[ix], equal))
	//}
	equal := strings.Contains(sub1.Operator, "=")
	operations = append(operations, GetOperator(sub1, equal))

	var newD Df

	for _, val := range d.Rows {
		keep := true
		//for col, f := range operations {
		//	if !f(&val, sub[col], true) {
		//		keep = false
		//		break
		//	}
		//}
		if !operations[0](&val, sub1, true) { //!SHIT
			keep = false //!SHIT
		} //!SHIT
		if keep {
			newD.Rows = append(newD.Rows, val)
		}
	}
	return newD
}

func main() {
	m := map[string]int{
		"abc":  1,
		"efgh": 2,
	}

	Row1 := Row{m, []float64{1.0, 2.0}}
	Row2 := Row{m, []float64{3.0, 5.0}}
	Row3 := Row{m, []float64{5.0, 7.0}}

	data := Df{[]Row{Row1, Row2, Row3}}

	fmt.Println(data.F("abc<5.0"))
}
