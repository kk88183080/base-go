package main

import "fmt"

func main() {

	_20test()
	fmt.Println("M")
}

func _20test() {
	defer fmt.Println("D")
	fmt.Println("F")
}
