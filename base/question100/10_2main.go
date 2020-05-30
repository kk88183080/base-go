package main

import "fmt"

func main() {
	test2()
}

//Go 中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的，所以编译错误
/*func test1() {
	a := [2]int{0, 1}
	b := [3]int{1, 2, 3}

	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("no equal")
	}
}
*/

func test() {
	a := [2]int{0, 1}
	b := [2]int{3, 4}

	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("no equal")
	}
}

func test2() {
	a := [2]int{0, 1}
	b := [2]int{0, 1}

	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("no equal")
	}
}
