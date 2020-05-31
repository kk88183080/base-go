package main

import "fmt"

func main() {
	//_23test()
	_23testslince()
}

func _23test() {

	if a := 1; false {

	} else if b := 2; false {

	} else {
		fmt.Println(a, b)
	}
}

func _23testslince() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:] // j-i 3-1=2 cap l-i 3-1=2
	fmt.Printf("len:%d, cap %d\n", len(s2), cap(s2))
	s2[1] = 4
	fmt.Println(s1) // 124
	fmt.Printf("len:%d, cap %d\n", len(s2), cap(s2))

	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1) // 124

	fmt.Printf("len:%d, cap %d\n", len(s2), cap(s2))
	fmt.Println(s2) // 2 4 5 6 7
}
