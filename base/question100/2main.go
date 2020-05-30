package main

import "fmt"

func main() {

	sli := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for i, v := range sli {
		m[i] = &v
	}

	for k, v := range m {
		fmt.Println(k, "-->", *v)
	}
}
