package main

import "fmt"

func main() {
	f := _63F(5)

	defer func() {
		fmt.Println(f()) //8
	}()

	defer fmt.Println(f()) //6

	i := f()
	fmt.Println(i) //7
}

func _63F(n int) func() int {

	return func() int {
		n++
		return n
	}
}
