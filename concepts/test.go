package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/apache/arrow/go/arrow"
// 	"github.com/apache/arrow/go/arrow/csv"
// )

func test() {

	// 	f, err := os.Open("C:/Users/kubistmi/projects/thesis/backend/cache/categories.csv")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	schema := arrow.NewSchema(
	// 		[]arrow.Field{
	// 			{Name: "category", Type: arrow.BinaryTypes.String},
	// 			{Name: "count", Type: arrow.PrimitiveTypes.Int64},
	// 		},
	// 		nil, // no metadata
	// 	)

	// 	rdr := csv.NewReader(f, schema, csv.WithChunk(-1))
	// 	defer rdr.Release()

	// 	n := 0
	// 	for rdr.Next() {
	// 		rec := rdr.Record()
	// 		for i, col := range rec.Columns() {
	// 			fmt.Printf("rec[%d][%q]: %v\n", n, rec.ColumnName(i), col)
	// 		}
	// 		n++
	// 	}
}
