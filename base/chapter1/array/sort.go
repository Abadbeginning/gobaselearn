package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{1, 9, 5}, {2, 3, 7}, {3, 6, 9}, {1, 8, 3}}
	// 按第二列排序
	firstIndex := 2
	result := ArraySort(nums, firstIndex-1)
	fmt.Println(result)
}

//按指定规则对numArray进行排序
func ArraySort(numArray [][]int, firstIndex int) [][]int {

	//检查
	if len(numArray) <= 1 {
		return numArray
	}

	if firstIndex < 0 || firstIndex > len(numArray[0])-1 {
		fmt.Println("Warning: Param firstIndex should between 0 and len(numArray)-1. The original array is returned.")
		return numArray
	}

	//排序
	mIntArray := &IntArray{numArray, firstIndex}
	sort.Sort(mIntArray)
	return mIntArray.mArr
}

type IntArray struct {
	mArr       [][]int
	firstIndex int
}

//IntArray实现sort.Interface接口
func (arr *IntArray) Len() int {
	return len(arr.mArr)
}

func (arr *IntArray) Swap(i, j int) {
	fmt.Println("交换的i, j", i, j)
	fmt.Println("交换前 --> ", arr.mArr[j], arr.mArr[i])
	arr.mArr[i], arr.mArr[j] = arr.mArr[j], arr.mArr[i]
	fmt.Println("交换后 --> ", arr.mArr[j], arr.mArr[i])
	fmt.Println()
}

func (arr *IntArray) Less(i, j int) bool {
	fmt.Println("i, j --> ", i, j)
	arr1 := arr.mArr[i]
	arr2 := arr.mArr[j]
	fmt.Println("arr1 --> ", arr1)
	fmt.Println("arr2 --> ", arr2)
	for index := arr.firstIndex; index < len(arr1); index++ {
		fmt.Println("arr1索引值 --> ", arr1[index])
		fmt.Println("arr2索引值 --> ", arr2[index])
		fmt.Println("结果为 --> ", arr1[index] < arr2[index])
		if arr1[index] < arr2[index] {
			return true
		} else if arr1[index] > arr2[index] {
			return false
		}
	}

	return i < j
}