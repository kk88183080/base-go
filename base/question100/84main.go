package main

import "fmt"

func main() {
	_84test()
}

func _84test() {

	a := [3]int{0, 1, 2}
	s := a[1:2] // [1]

	s[0] = 11         // [11] [0, 11, 2] s len: 1, cap:2
	s = append(s, 12) // [11, 12]
	s = append(s, 13) // [11, 12, 13]
	s[0] = 21         // [21, 12, 13]

	fmt.Println(a) // [0, 11, 12]
	fmt.Println(s) // 21, 12, 13

}
