package tab

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/kubistmi/goframe/vec"
)

// ReadCsv ...
func ReadCsv(reader io.Reader) Table {
	testRows := 100
	file := csv.NewReader(reader)

	file.ReuseRecord = false
	colnames, errH := file.Read()
	if errH != nil {
		return Table{
			err: fmt.Errorf("Error loading data at line: %v", 1),
		}
	}

	parseTest := make([][]string, testRows)
	var err error
	j := 0
	end := false
	for i := 0; i < testRows; i++ {
		parseTest[i], err = file.Read()
		if err == io.EOF {
			end = true
			break
		} else if err != nil {
			return Table{
				err: fmt.Errorf("Error loading data at line: %v", j+2),
			}
		}
		j++
	}

	coltypes := make(map[string]string, len(colnames))

	for col, name := range colnames {
		for row := 0; row < j; row++ {
			guess := testType(parseTest[row][col])
			if guess == "str" {
				coltypes[name] = "str"
				break
			} else if guess == "int" {
				_, ok := coltypes[name]
				if !ok {
					coltypes[name] = "int"
				} else if coltypes[name] == "int" {
					continue
				} else {
					coltypes[name] = "string"
					break
				}
			}
		}
	}

	strCols := make([]string, 0, len(colnames))
	intCols := make([]string, 0, len(colnames))

	for ix, val := range coltypes {
		if val == "str" {
			strCols = append(strCols, ix)
		} else {
			intCols = append(intCols, ix)
		}
	}

	strings := make(map[string][]string, len(strCols))
	ints := make(map[string][]int, len(intCols))

	for _, val := range strCols {
		strings[val] = make([]string, 1000000)
	}

	for _, val := range intCols {
		ints[val] = make([]int, 1000000)
	}

	i := 0
	for row := 0; row < j; row++ {
		for col, name := range colnames {
			if coltypes[name] == "str" {
				strings[name][i] = parseTest[row][col]
			} else {
				ints[name][i], _ = strconv.Atoi(parseTest[row][col])
			}
		}
		i++
	}

	if !end {
		for {
			row, err := file.Read()
			i++
			if err == io.EOF {
				i--
				break
			} else if err != nil {
				return Table{
					err: fmt.Errorf("Error loading data at line: %v", i+2+testRows),
				}
			}
			for col, name := range colnames {
				if coltypes[name] == "str" {
					strings[name][i] = row[col]
				} else {
					ints[name][i], err = strconv.Atoi(row[col])
					if err != nil {
						return Table{
							err: fmt.Errorf("Error loading data for column %s at line: %v", name, i+2+testRows),
						}
					}
				}
			}
		}
	}

	data := make(map[string]vec.Vector, len(colnames))

	for _, name := range colnames {
		data[name] = prepareVec(name, i, ints, strings)
	}

	df, err := NewDf(data)
	if err != nil {
		return Table{
			err: fmt.Errorf("Error in final compilation of Table"),
		}
	}
	return df
}

func testType(input string) string {

	_, err := strconv.Atoi(input)
	if err != nil {
		return "str"
	}
	return "int"

}

func prepareVec(name string, i int, ints map[string][]int, strings map[string][]string) vec.Vector {
	if val, ok := ints[name]; ok {
		data := make([]int, i+1)
		copy(data, val[:i+2])
		return vec.NewVec(data, nil)
	}
	data := make([]string, i+1)
	copy(data, strings[name][:i+2])
	return vec.NewVec(data, nil)
}
