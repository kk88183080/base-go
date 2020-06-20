package main

import "fmt"

func main() {

	arr := [5]int{23, 0, 12, 56, 34}
	insertSort(&arr)
}

/**
	排序的过程
第一步 把23当成是一个有序的数组
第二步 给0找个位置，进行插入
第三步
   把0保存到一个变量中，
   把插入的位置保存到一个变量中
第4步
   从有序数据的最后的一个数进行比较
第5步
   在找的过程中，有序数组元素要后移
*/

/**
	过程==>23, 0, 12, 34, 56 原数据
1> [23 ，0] 12 56, 34
	12
2> [23 , 0, 0] 56, 34
   [23, 12, 0] 56, 34
3> [23, 12, 0, 0], 34
   [23, 12, 12, 0], 34
   [23, 23, 12, 0], 34
   [56, 23, 12, 0], 34
*/

/**
插入排序,从大到小排序
*/
func insertSort(arr *[5]int) {

	for j := 1; j < len(arr); j++ {
		insertVal := arr[j]
		insertIndex := j - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 这个值后移
			insertIndex--
		}
		if insertIndex+1 != j {
			arr[insertIndex+1] = insertVal
			fmt.Printf("内部%d次, %v\n", j, *arr)
		}

		fmt.Printf("%d 次， %v\n", j, *arr)
	}

}
