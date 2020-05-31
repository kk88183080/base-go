package main

import "fmt"

func main() {
	//_14Test()
	//_14Test1()
	_14test2()
}

/*func _14Test() {
	str := "hello"
	// 编译出错，字符串是只读的
	str[0] = 'x'
	fmt.Println(str)
}*/

func _14Test1() {
	p := 1
	_14incr(&p)
	fmt.Println(p)
}

func _14incr(i *int) int {
	*i++
	return *i
}

func _14test2() {
	_14AddTest()
	_14AddTest(1, 2)
	_14AddTest(1, 2, 3)
	_14AddTest([]int{1, 2, 3}...)
	// 可变参数不能接受切片，但可能接受切片里面的值
	//_14AddTest([]int{1,2})
}

func _14AddTest(args ...int) {
	sum := 0
	for _, v := range args {
		sum += v

	}

	fmt.Println(sum)
}
