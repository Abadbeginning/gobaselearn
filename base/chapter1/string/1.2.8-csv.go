package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.csv")
	if err != nil {
		panic(err)
	}

    defer file.Close()

    reader := csv.NewReader(file)
    // reader.Comma = '#'

    for {
        record, error := reader.Read()
        fmt.Printf("%v\n", record)
        if error != nil {
            fmt.Println(error)
            break
        }
        
    }
}
