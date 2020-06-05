package main

import "fmt"

func main() {
	//_61test()
	_61test2()
}

func _61test() {

	var k = 1

	var s = []int{1, 2}
	k, s[k] = 0, 3           // k=0, s=1,3
	fmt.Println(s[0] + s[1]) //4
}

func _61test2() {
	var k = 9
	for k = range []int{} {
	}

	fmt.Println(k) // 9

	for k = 0; k < 3; k++ {

	}
	fmt.Println(k) //3

	for k = range (*[3]int)(nil) {
	}
	fmt.Println(k) //2
}
