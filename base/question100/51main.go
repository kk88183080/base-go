package main

import "fmt"

func main() {
	_51test2()
}

/*func _51test()  {
	var f  func()
	var f1  func()

	if f == f1 {
		fmt.Println("f is not equals f1")
	}
}*/
// 编译错误，函数只能与nil值进行比较

type _51T struct {
	i int
}

/*func _51Test1()  {
	m:= make(map[int]_51T)
	m[0].i = 10

	fmt.Println(m[0].i)
}*/
// 编译错误

func _51test2() {
	m := make(map[int]_51T)

	m[0] = _51T{1}
	fmt.Println(m[0].i)
}
