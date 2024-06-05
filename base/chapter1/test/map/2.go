package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// fmt.Println("Hello, 世界")
	// str1()
	// str4()
	// str5()
	// str6()
	str7()
}

func str1() {
	name := "萨克斯是"
	// fmt.Println(len(name))
	// fmt.Println([]byte(name))
	// fmt.Println(name[0], name[1], name[2])
	// fmt.Println(name[:3])

	other := []rune(name)
	fmt.Println(other)
	fmt.Println(string(other[0]))
	fmt.Println(string(other[:2]))
}

func str2() {
	// >>> s = "憨pi"
	// # 采用 utf-8 编码（等价于 Go 的 []byte 数组）
	// # "憨" 需要 230 134 168 三个整数来表示
	// # 而 "p" 和 "i" 均只需 1 个字节，分别为 112 和 105
	// [c for c in s.encode("utf-8")]
	// [230, 134, 168, 112, 105]

	// # 采用 unicode 编码（类似于 Go 的 [] rune数组）
	// # 所有字符都只需要1个整数表示
	// # 但对于 ASCII 字符而言，不管什么编码，对应的数值不变
	// >>> [ord(c) for c in s]
	// [25000, 112, 105]
	// s := "憨pi"
	s := "憨pi"
	fmt.Println([]byte(s)) // [230 134 168 112 105]

	fmt.Println([]rune(s)) // [25000 112 105]
}

func str3() {
	s1 := []byte{230, 134, 168, 112, 105}
	fmt.Println(string(s1)) // 憨pi

	s2 := []rune{25000, 112, 105}
	fmt.Println(string(s2)) // 憨pi
}

func str4() {
	s1 := []int8{1, 2, 3, 4}
	s2 := *(*[]int16)(unsafe.Pointer(&s1))
	fmt.Println(s2)
	fmt.Println(2<<8 + 1)
	fmt.Println(4<<8 + 3)
}

func str5() {
	str := "abc"
	slice := *(*[]byte)(unsafe.Pointer(&str))
	fmt.Println(slice)      // [97 98 99]
	fmt.Println(cap(slice)) // 10036576

}

func str6() {
	slice := []byte{97, 98, 99}
	str := *(*string)(unsafe.Pointer(&slice))
	fmt.Println(str)  // abc
	slice[0] = 'A'
	fmt.Println(str)  // Abc
}

func str7() {
	fmt.Println(StringToBytes("abc")) // [97 98 99]
	fmt.Println(cap(StringToBytes("abc")))
	fmt.Println(BytesToString([]byte{97, 98, 99})) // abc
}

func StringToBytes(s string) []byte {
	// 既然字符串转切片，会丢失容量
	// 那么加上去就好了，做法也很简单
	// 新建一个结构体，将容量（等于长度）加进去
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	// 切片转字符串就简单了，直接转即可
	// 转的过程中，切片的 Cap 字段会丢弃
	return *(*string)(unsafe.Pointer(&b))
}

