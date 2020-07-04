package main

import (
	"fmt"
	"strings"
)

func main() {
	//arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println("max val:", SelectSort(arr))
	//fmt.Println("main:", arr)

	//arr1 := []string{"1", "9", "2", "8", "3", "7", "4", "6", "5", "10"}
	//SelectSortString(arr1)
	//fmt.Println("main:", arr1)

	//arr1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k"}
	//SelectSortString(arr1)
	//fmt.Println("main:", arr1)

	//arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//InsertSort(arr)
	//fmt.Println("main:", arr)

	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	MaoPaoSort(arr)
	fmt.Println("main:", arr)
}

/**
选择排序
*/
func SelectSort(arr []int) int {

	fmt.Println("SelectSort source:", arr)

	length := len(arr)

	if length <= 1 {
		return arr[0]
	}

	for i := 0; i < length-1; i++ {
		max := i
		for j := i + 1; j < length; j++ {
			if arr[max] < arr[j] {
				max = j
			}
		}
		if max != i {
			arr[max], arr[i] = arr[i], arr[max]
		}
	}

	fmt.Println("SelectSort sorted:", arr)

	return arr[0]
}

/**
字符串排序
*/
func SelectSortString(arr []string) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		index := i
		for j := i + 1; j < length; j++ {

			if strings.Compare(arr[index], arr[j]) < 0 { // 从大到小排序
				index = j
			}
		}

		if index != i {
			arr[index], arr[i] = arr[i], arr[index]
		}

	}
}

/**
插入排序，左边插入排序，右边插入排序
1,2,3
*/
func InsertSort(arr []int) {
	length := len(arr)

	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {

		backup := arr[i]
		j := i - 1
		for j >= 0 && backup > arr[j] {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = backup
	}
}

/**
冒泡排序
*/
func MaoPaoSort(arr []int) {

	length := len(arr)

	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
