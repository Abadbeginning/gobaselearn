package main

import (
	"fmt"
	"unsafe"
)

func main() {
  // newMap1()
  // map2()
  // map3()
  // map4()
  // arr1()
  // arr2()
  // arr3()
  arr4()
}

func newMap1() {
  // var s1 = []int{1, 7:2, 3}
  // fmt.Println(s1)
  
  // new 函数接收一个类型，创建对应的零值，然后返回其指针
  var s = new([]int)
  // 所以使用这种方式只会创建切片本身，但是切片对应的底层数组是不会创建的
  // 内部的指针是一个 nil、长度和容量都是 0，但使用 append 的话会自动创建底层数组
  *s = append(*s, 1, 2, 3, 4)
  fmt.Println(s)   // &[1 2 3 4]
  fmt.Println(*s)  // [1 2 3 4]
}

func arr4() {
  var s1 = []int{1, 2, 3, 4, 5}
  var s2 = []int{6, 7, 8}
  // 将 s1 拷贝到 s2 中，会从头开始拷贝
  copy(s2, s1)
  // 因为 s2 长度为 3，因此只会拷贝 3 个
  fmt.Println(s2)  // [1 2 3]

  var s3 = []int{1, 2, 3}
  var s4 = []int{4, 5, 6, 7, 8}
  // 将 s3 拷贝到 s4 中
  copy(s4, s3)  
  fmt.Println(s4)  // [1 2 3 7 8]

  var s5 = []int{1, 2, 3, 4, 5}
  var s6 = make([]int, 1, 3)
  copy(s6, s5)
  // 切片拷贝前后的长度和容量都是不变的
  // 并且能往切片里面拷贝多少个元素，取决于切片的长度、而不是容量
  fmt.Println(s6)  // [1]
  fmt.Println(s6[: 3]) // [1 0 0]

  var s7 = []int{1, 2, 3}
  var s8 = []int{3, 4, 5}
  // copy 的行为相当于覆盖，如果想追加呢？
  s7 = append(s7, s8[1:]...)
  fmt.Println(s7)  // [1 2 3 4 5]
}

func arr3() {
  s1 := []int{1, 2, 3}
  s2 := s1
  s2[0] = 666
  fmt.Println(s1)  // [666 2 3]
  fmt.Println(s2)  // [666 2 3]
}

func arr2() {
  var arr = []int{1, 2, 3}
  s1 := arr[1:]
  // 写成 s2 = s1[:] 或者 s2 = s1 也可以
  s2 := arr[1:]
  fmt.Println(s1, s2) // [2 3] [2 3]
  // 因为是同一个数组，所以地址一样
  fmt.Println(&s1[0], &s2[0]) // 0xc0000b2020 0xc0000b2020

  // 此时 s1 和 s2 都是 [2, 3]
  // 下面给 s2 扩容
  s2 = append(s2, 4)
  // 地址不一样了
  fmt.Println(&s1[0], &s2[0]) // 0xc0000b2020 0xc0000ae040
}

func arr1() {
  var s = make([]int, 0, 3)
  s = append(s, 1)
  fmt.Printf("%p\n", &s[0]) // 0xc00000c150
  s = append(s, 2)
  fmt.Printf("%p\n", &s[0]) // 0xc00000c150
  s = append(s, 3)
  fmt.Printf("%p\n", &s[0]) // 0xc00000c150

  // 如果再 append，那么容量肯定不够了
  s = append(s, 4)
  fmt.Printf("%p\n", &s[0]) // 0xc00000a360
}

func map4() {
  var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
  s1 := arr[1: 3]
  s2 := s1[3: 6]
  fmt.Println(s1) // [b c]
  fmt.Println(s2) // [e f g]

  // [:] 这种方式只会获取当前切片可以看到的元素
  // 换句话说可以看到的元素的个数等于切片长度
  fmt.Println(s2[:])  // [e f g]
  fmt.Println(s2[: 4]) // [e f g h]
  fmt.Println(s2[: cap(s2)])
  fmt.Println(s2[4])
}

func map3() {
  var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
  // 如果是 s[m: n]，那么等价于 s[m: n: len(arr)]
  // 默认可以向后扩展到数组的结束位置，容量为 len(arr) - m
  // 这里是 s[1: 3: 5]，容量为 5 - 1
  // 表示可以向后扩展到数组长度为 5 的位置
  s := arr[1: 3: 5]
  fmt.Println("计算s --> ",  cap(s))
  fmt.Println(s[2: 4])
  fmt.Println("----------我是分界线----------")
  fmt.Println(s[3: 5])
  /*
     [d e]
     ----------我是分界线----------
     panic: runtime error: slice bounds out of range [:5] with capacity 4
  */
}

func map2() {
  var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}

  s := arr[1: 3]
  fmt.Println(s)
  fmt.Println(s[3: 7])
  fmt.Println("----------我是分界线----------")
  fmt.Println(s[3])
}

func map1() {
  // [] 里面什么都不写的话, 表示创建一个切片
  var s1 = []int{1, 2, 3}
  var s2 = []string{"刷到过1", "2", "3"}
  var s3 = []float64{1.1, 2.2, 3.3}

  // 查看变量所占内存的话, 可以使用 unsafe.Sizeof
  fmt.Println(unsafe.Sizeof(s1))  // 24
  fmt.Println(unsafe.Sizeof(s2))  // 24
  fmt.Println(unsafe.Sizeof(s2[0]))  // 24
  fmt.Println(unsafe.Sizeof(s3))  // 24
}