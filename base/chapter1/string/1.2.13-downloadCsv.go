package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// fmt.Println("Hello, 世界")
	// http.Handle("/down", Welcome)
	// Welcome()
	http.HandleFunc("/down", Welcome)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



func Welcome(w http.ResponseWriter, r *http.Request) {
	filename := "exportusers.csv"

	column := [][]string {
		{"id", "name", "email"},
		{"1", "张三", "rdlnk@example.com"},
		{"2", "李四", "xcvkp@example.com"},
		{"3", "王五", "rdlnk@example.com"},
	}
	GenerateCsv(filename, column)

	// 读取文件
	file, err := os.Open(filename)
	content, err := io.ReadAll(file)

	w.Header().Set("Conent-Type", "Apllication/octet-stream")
	w.Header().Set("Cotent-Disposition", "attachment; filename=\""+filename+"\"")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		w.Write(content)
	}
	
}

func GenerateCsv(fileName string, data [][]string) {
	fp, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("创建文件["+fileName+"]句柄失败，%v", err)
		return
	}

	defer fp.Close()

	//  写入utf8 bom
	fp.WriteString(string([]byte{0xEF, 0xBB, 0xBF}))
	w := csv.NewWriter(fp)
	w.WriteAll(data)
	w.Flush()
}