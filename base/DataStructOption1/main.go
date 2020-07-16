package main

import (
	"fmt"
	"math/rand"
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

	//arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println("main:", QuickSort1(arr))

	//arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//JiOusort(arr)
	//fmt.Println("main:", arr)

	//arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println("main:", GuiBingSort(arr))

	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	XiErSort(arr)
	fmt.Println("main:", arr)
}

/**
希尔排序
*/
func XiErSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	gap := length / 2
	for gap > 0 {
		for i := 0; i < gap; i++ {
			XiErSortStep(arr, i, gap)
		}
		gap /= 2
	}
}

func XiErSortStep(arr []int, start int, gap int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := start + gap; i < length; i += gap { //插入排序

		backup := arr[i] // 备份数据
		j := i - gap

		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}

		arr[j+gap] = backup
	}
}

/**
归并排序
1, 9, 2, 8, 3, 7, 4, 6, 5, 10
第一轮，1, 9, 2, 8, 3, 7  			4, 6, 5, 10
第二轮  1, 9, 2  		8, 3, 7
分段排序，然后再合并
*/
func GuiBingSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	if length == 10 { // 优化，使用其他方法

	}
	mid := length / 2
	leftarr := GuiBingSort(arr[:mid])
	rightarr := GuiBingSort(arr[mid:])

	return Merge(leftarr, rightarr)
}

/**
归并排序
合并数据
*/
func Merge(leftarr []int, rightarr []int) []int {
	leftindex := 0
	rightindex := 0

	lastarr := []int{}
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] { // 小于
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		} else if leftarr[leftindex] > rightarr[rightindex] { // 大于
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		} else { // 相等
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		}
	}

	for leftindex < len(leftarr) { // 把左边没有归并完的数据进行归并
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}

	for rightindex < len(rightarr) { // 把右边没有归并完的数据进行归并
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}

	return lastarr
}

/**
奇偶排序
1,7,3,8,2,9,5,6,4,0

*/
func JiOusort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	isSorted := false

	for isSorted == false {

		isSorted = true
		for i := 1; i < length-1; i += 2 { // 奇数位
			if i+1 <= length-1 && arr[i] > arr[i+1] { //是不要交换
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}

		for i := 0; i < length-1; i += 2 { // 偶数位
			if i+1 <= length-1 && arr[i] > arr[i+1] { //是不要交换
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}

		fmt.Println("con:", arr)
	}
}

/**
快速排序优化
*/
func QuickSort1(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	//n := length - 1           // 这个位置值要>=0 <=length-1
	n := rand.Int() % length  // 获取随机数
	splitdata := arr[n]       // 以第一个为基准
	low := make([]int, 0, 0)  // 存储比基准小的
	high := make([]int, 0, 0) // 存储比基准大的
	mid := make([]int, 0, 0)  // 存储与基准相等的
	mid = append(mid, splitdata)

	// 第一个数不用处理
	for i := 0; i < length; i++ {
		if n == i {
			continue
		} else {
			if arr[i] < splitdata { // 存储小的
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
	}

	low, high = QuickSort(low), QuickSort(high) // 切割，递归处理

	rs := append(append(low, mid...), high...)
	return rs

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
