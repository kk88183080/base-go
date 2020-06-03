package main

import "fmt"

func main() {
	//_50test()
	//_50test2()
	//_50test3()
	_50test4()
}

type _50T struct {
	s []int
}

func _50foo(t []int) {
	t[0] = 100
}

func _50test() {
	var t = _50T{
		s: []int{1, 2, 3},
	}

	_50foo(t.s)

	fmt.Println(t)

}

func _50test2() {
	var a = func(i int) bool {
		switch i {
		case 1:
		case 2:
			return true
		}

		return false
	}

	fmt.Println(a(1))
	fmt.Println(a(2))

}

//Go 语言的 switch 语句虽然没有"break"，但如果 case 完成程序会默认 break，
//可以在 case 语句后面加上关键字 fallthrough，这样就会接着走下一个 case 语句（不用匹配后续条件表达式）。
//或者，利用 case 可以匹配多个值的特性。

func _50test3() {
	var a = func(i int) bool {
		switch i {
		case 1:
			fmt.Println("1 start")
			fallthrough
			//fmt.Println("1 end")
		case 2:
			return true
		}

		fmt.Println("return val")
		return false
	}

	fmt.Println(a(1)) // true
	fmt.Println(a(2)) // true
}

func _50test4() {
	var a = func(i int) bool {
		switch i {
		case 1, 2:
			return true
		}
		return false
	}

	fmt.Println(a(1))
	fmt.Println(a(2))
}
