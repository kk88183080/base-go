package main

import "fmt"

func main() {
	//_66test()
	//_66test2()
	//_67test()
	_67test1()
}

type _66T struct {
	n int
}

func _66test() {

	ts := [2]_66T{}
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ") //0
		}
	}

	fmt.Print(ts) //[{0},{9}]
}

//for-range 循环数组。此时使用的是数组 ts 的副本，所以 t.n = 3 的赋值操作不会影响原数组。

func _66test2() {
	ts := [2]_66T{}

	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ") // 9
		}
	}

	fmt.Print(ts) // [{0},{9}]
}

//for-range 数组指针。for-range 循环中的循环变量 t 是原数组元素的副本。如果数组元素是结构体值，则副本的字段和原数组字段是两个不同的值。

func _67test() {

	ts := [2]_66T{}

	for i := range ts[:] {
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Print(ts[i].n, " ") //9
		}
	}

	fmt.Print(ts) //[{0},{9}]
}

//for-range 切片。for-range 切片时使用的是切片的副本，但不会复制底层数组，换句话说，此副本切片与原数组共享底层数组。

func _67test1() {

	ts := [2]_66T{}

	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ") //9

		}
	}

	fmt.Print(ts) // [{3},{9}]
}
