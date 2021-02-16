package tab

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/kubistmi/goframe/vec"
)

func (df Table) printRecord(row int) string {
	return ""
}

func nchar(s string) int {
	return utf8.RuneCountInString(s)
}

func lim(max bool, x ...int) int {
	m := x[0]
	if max {
		for _, v := range x {
			if v > m {
				m = v
			}
		}
	} else {
		for _, v := range x {
			if v < m {
				m = v
			}
		}
	}

	return m
}

func pad(s string, n int) string {
	nc := nchar(s)

	var sb strings.Builder
	if nc >= n {
		sb.WriteString(" ")
		srun := []rune(s)[:(n - 5)]
		sb.WriteString(string(srun))
		// sb.WriteString(s[:(n - 5)])
		sb.WriteString("... ")
	} else {
		numspace := n - nc
		right := numspace / 2
		left := numspace - right
		for i := 0; i < left; i++ {
			sb.WriteString(" ")
		}
		sb.WriteString(s)
		for i := 0; i < right; i++ {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

// Print ...
func (df Table) Print() string {
	colmax := 20
	dfV := df.Head(10)
	colLen := make(map[string]int, len(df.names))
	strVecs := make(map[string][]string, len(dfV.names))
	strNas := make(map[string]vec.Set, len(dfV.names))
	var sb strings.Builder
	var sepb strings.Builder

	for _, val := range df.names {
		colLen[val] = lim(true, lim(false, nchar(val), colmax), 10)
		strVecs[val], strNas[val] = dfV.data[val].ToStr().Get()
	}

	for _, col := range df.names {
		for i := 0; i < 10; i++ {
			colLen[col] = lim(true, colLen[col], lim(false, nchar(strVecs[col][i]), colmax))
		}
	}

	width := 3
	for _, l := range colLen {
		width = width + l
	}
	for i := 0; i < width; i++ {
		sepb.WriteString("-")
	}
	sepb.WriteString("\n")

	sb.WriteString(sepb.String())
	for _, col := range df.names {
		sb.WriteString("|")
		sb.WriteString(pad(col, colLen[col]))
		sb.WriteString("|")
	}
	sb.WriteString("\n")
	sb.WriteString(sepb.String())

	for i := 0; i < 10; i++ {
		for _, col := range df.names {
			sb.WriteString("|")
			sb.WriteString(pad(strVecs[col][i], colLen[col]))
			sb.WriteString("|")
			fmt.Println()
		}
		sb.WriteString("\n")
	}

	sb.WriteString(sepb.String())
	return sb.String()
}

func (df Table) String() string {
	return df.Print()
}
