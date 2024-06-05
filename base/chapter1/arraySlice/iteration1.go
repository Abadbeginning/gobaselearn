package main

import "fmt"

func main() {
	// fmt.Println("Hello, 世界")
	// iter1()
	// iter2()
	iter3()
}

func iter1() {
	array := []int{11, 22, 33, 44, 55}
	for k, v := range array {
		fmt.Println(k, v, &array[k])
	}

	// 忽略索引
	for _, v := range array {
		fmt.Println(v)
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func iter2() {
	seq := []string{"are", "you", "ok", "today"}
	index := 2
	fmt.Println(seq)
	fmt.Println(seq[:index], seq[index+1:])
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}

func iter3() {
	array := make([]string, 5)
	array[0] = "are you ok"
	string := array[0]
	fmt.Println(string)
}

func arrayToString(arr []string) string {
	var result string
	for _, v := range arr {
		result += v
	}

	return result
}