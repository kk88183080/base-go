package main

import "fmt"

func main() {

	fmt.Println(_18f1())
	fmt.Println(_18f2())
	fmt.Println(_18f3())
}

func _18f1() (r int) {
	defer func() {
		r++
	}()

	return 0
}

func _18f2() (r int) {

	t := 5
	defer func() {
		t = t + 5
	}()

	return t
}

func _18f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)

	return 1
}
