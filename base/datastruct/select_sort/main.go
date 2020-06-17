package main

import "fmt"

func main() {
	selectSort()
}

/**
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复第二步，直到所有元素均排序完毕。
*/
// 从小到大排序
func selectSort() {
	array := [5]int{5, 4, 3, 2, 1}
	length := len(array)

	fmt.Println(array)
	for i := 0; i < length-1; i++ {
		minindex := i
		for j := i + 1; j < length; j++ {
			if array[minindex] > array[j] { // 不做地址更新，进行性能优化
				minindex = j
			}
		}
		if minindex != i {
			array[i], array[minindex] = array[minindex], array[i]
			fmt.Println(array)
		}
	}
	fmt.Println(array)
	// 1, 4, 3, 2, 5
	// 1, 2, 3, 4, 5
}
