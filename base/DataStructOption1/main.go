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

	//arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	////MaoPaoSort(arr)
	//MaoPaoSort2(arr)
	//fmt.Println("main:", arr)

	/*	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
		//MaoPaoSort(arr)
		//HeapSortMax(arr, len(arr))
		HeapSort(arr)
		fmt.Println("main:", arr)*/

	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println("main:", QuickSort(arr))
}

/**
快速排序
*/
func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	splitdata := arr[0]       // 以第一个为基准
	low := make([]int, 0, 0)  // 存储比基准小的
	high := make([]int, 0, 0) // 存储比基准大的
	mid := make([]int, 0, 0)  // 存储与基准相等的
	mid = append(mid, splitdata)

	// 第一个数不用处理
	for i := 1; i < length; i++ {
		if arr[i] < splitdata { // 存储小的
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, high = QuickSort(low), QuickSort(high) // 切割，递归处理

	rs := append(append(low, mid...), high...)
	return rs

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

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

/**
冒泡排序2
// 1, 2, 3

// 2, 1, 3
// 2, 3, 1

//

*/
func MaoPaoSort2(arr []int) {

	length := len(arr)

	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		isSwap := false
		for j := 0; j < length-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

/**
堆排序
1, 2, 3, 4, 5, 6, 7
		1
	2 		3
4    5    6  7
*/
func HeapSortMax(arr []int, length int) {

	//length := len(arr)

	if length <= 1 {
		return
	}

	depth := length/2 - 1 // 树的深度，左节点2*n+1，右节点2*n-1
	for i := depth; i >= 0; i-- {
		topmax := i
		leftChild := 2*i + 1
		rightChild := 2*i + 2

		if leftChild <= length-1 && arr[leftChild] > arr[topmax] {
			topmax = leftChild
		}

		if rightChild <= length-1 && arr[rightChild] > arr[topmax] {
			topmax = rightChild
		}

		if topmax != i { // 完成交换
			arr[topmax], arr[i] = arr[i], arr[topmax]
		}
		fmt.Println(arr)
	}

}

func HeapSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		lastmesslen := length - i
		HeapSortMax(arr, lastmesslen)
		if i < length {
			arr[0], arr[lastmesslen-1] = arr[lastmesslen-1], arr[0]
		}
	}

}
