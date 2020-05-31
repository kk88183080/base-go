package main

import "fmt"

func main() {

	//fmt.Println(_30test(3))
	//_30testArray()
	_30testArray1()
	//_30testSlice()
}

func _30test(n int) (r int) {

	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f() // 6

	f = func() {
		r += 2
	}

	return n + 1
}

// 这个说明 range 时数组是值复制
func _30testArray() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[i] = v
	}
	fmt.Println(a)
	fmt.Println(r)
}

func _30testArray1() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[i] = v
	}
	fmt.Println(a)
	fmt.Println(r)
}

// 这个说明 range 时切片是地址引用
func _30testSlice() {
	var a = []int{1, 2, 3, 4, 5}
	r := [5]int{}

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[i] = v
	}
	fmt.Println(a)
	fmt.Println(r)
}
