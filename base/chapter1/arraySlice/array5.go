package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// fmt.Println("Hello, 世界")
	str := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	a, _ := RandomInt(str, len(str)-2)
	fmt.Println(a)
}

func RandomInt(strings []string, length int) (string, error) {
	if length <= 0 {
		return "", errors.New("字符串的长度不能小于0")
	}

	if length >= len(strings) || length <= 0 {
		return "", errors.New("参数长度非法")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < length; i++ {
		str += strings[i]
	}
	return str, nil
}
