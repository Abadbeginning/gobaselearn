package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取6位随机字符串
	b := GetRandomString(6)
	fmt.Println(b)

	// fmt.Println("Hello World")
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
    fmt.Println(result)
	return string(result)
}
