package main

import "fmt"

func main() {
	_69test1()
}

func _69test(x int) (func(), func()) {

	return func() {
			fmt.Println(x)
			x += 10
		}, func() {
			fmt.Println(x)
		}
}

func _69test1() {
	a, b := _69test(100)
	a()
	b()
}

// 100 110
