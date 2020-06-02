package main

import "fmt"

func main() {
	_47test1()
}

// 编译错误
/*func _47test()  {
	data := []int{1, 2, 3}
	i:=0
	++i
	fmt.Println(data[i++])
}*/

// 修复如下
func _47test() {
	data := []int{1, 2, 3}
	i := 0
	i++
	fmt.Println(data[i])
}

func _47test1() {
	x := 1
	fmt.Println(x)

	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}

	fmt.Println(x)
}

// 变量隐藏。使用变量简短声明符号 := 时，如果符号左边有多个变量，
// 只需要保证至少有一个变量是新声明的，并对已定义的变量尽进行赋值操作。但如果出现作用域之后，就会导致变量隐藏的问题，就像这个例子一样。
