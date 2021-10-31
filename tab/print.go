package tab

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/kubistmi/goframe/vec"
	"github.com/olekukonko/tablewriter"
)

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
func (df Table) PrintOwn() string {
	n := lim(false, 10, df.size[0])
	colmax := 60
	dfV := df.Head(n)
	colLen := make(map[string]int, len(df.names))
	strVecs := make(map[string][]string, len(dfV.names))
	strNas := make(map[string]vec.NA, len(dfV.names))
	var sb strings.Builder
	var sepb strings.Builder
	var err error

	for _, val := range df.names {
		colLen[val] = lim(true, lim(false, nchar(val)+5, colmax), 10)
		strVecs[val], strNas[val], err = dfV.data[val].ToStr().Get()
		if err != nil {
			return (fmt.Errorf("error in accessing the data").Error())
		}
	}

	for _, col := range df.names {
		for i := 0; i < n; i++ {
			colLen[col] = lim(true, colLen[col], lim(false, nchar(strVecs[col][i]), colmax))
		}
	}

	width := 0
	for _, l := range colLen {
		width = width + l + 2
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

	for i := 0; i < n; i++ {
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
	return df.PrintOwn()
}

// based on https://github.com/olekukonko/tablewriter
// TODO: how to make it comply with Stringer()?
func (df Table) Print() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(df.names)

	df = df.Rows([]int{0, 1, 2, 3, 4})
	textmat := make([][]string, df.size[0])
	for ix := range textmat {
		textmat[ix] = make([]string, df.size[1])
	}

	for i, col := range df.names {
		vals, NA, err := df.Pull(col).ToStr().Get()
		if err != nil {
			panic("wtf!")
		}
		for j, v := range vals {
			if NA.Get(j) {
				textmat[j][i] = "NA"
			} else {
				textmat[j][i] = v
			}
		}
	}

	for _, val := range textmat {
		table.Append(val)
	}
	table.Render()

}
