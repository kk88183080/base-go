package main

import "fmt"

func main() {
	i := []int{1, 2, 3}
	_9hello(i...)
	fmt.Println(i[0])
}

func _9hello(num ...int) {
	num[0] = 18
}
