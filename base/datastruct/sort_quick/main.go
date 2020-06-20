package main

import "fmt"

func main() {

	arr := [18]int{-9, 78, 0, 23, -567, 70, -9, 78, 0, 23, -567, 70, -9, 78, 0, 23, -567, 70}
	quickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}

func quickSort(left int, right int, arr *[18]int) {

	l := left
	r := right

	// 获取到中轴
	pivot := arr[(left+right)/2]
	temp := 0

	for l < r { // 目标，比pivot小的放在左边，大的放在右边

		for arr[l] < pivot {
			l++
		}

		for arr[r] > pivot {
			r--
		}

		if l >= r { // 比较到最中间了
			break
		}

		temp = arr[r]
		arr[r] = arr[l]
		arr[l] = temp
		if arr[l] == pivot {
			r--
		}

		if arr[r] == pivot {
			l++
		}
	}

	if l == r {
		l++
		r--
	}

	if left < r {
		quickSort(left, r, arr)
	}

	if right > l {
		quickSort(l, right, arr)
	}
}
