package main

import "fmt"

func main() {
	//_45test2()
	_45test4()

}

// 编译出错
//func _45test()  {
//	one:=0
//	one:=1
//}

func _45test1() {
	one := 0
	one, tow := 1, 2

	fmt.Println(one, tow)
}

func _45test2() {
	x := []int{
		1,
		2, // 多行时，这块得加，号
	}
	_ = x

	y := []int{1, 2} // 一行时，最后一个元素可以不加，号
	_ = y

	z := []int{1, 2}
	z = z
	fmt.Println(z)
}

func _45test3(x byte) {
	fmt.Println(x)
}

func _45test4() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	_45test3(c)
}

//与 rune 是 int32 的别名一样，byte 是 uint8 的别名，别名类型无序转换，可直接转换。
