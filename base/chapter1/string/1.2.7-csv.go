package main

import (
	"encoding/csv"
	// "fmt"
	"os"
)

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 写入utf-8 BOM
	f.WriteString("\xEF\xBB\xBF")

	// NewWriter()函数返回一个Writer对象
	w := csv.NewWriter(f)

	w.Write([]string{"学号","姓名","分数"})
	w.Write([]string{"1", "张三", "100"})
	w.Write([]string{"2", "张四", "90"})
	w.Write([]string{"3", "李四", "80"})
	w.Write([]string{"4", "李渊", "70"})
	w.Flush()
}
