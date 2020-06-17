package main

import "fmt"

func main() {
	maopao()
}

func maopao() {
	// 大到小排
	array := [5]int{10, 22, 8, 60, 4}
	length := len(array)
	fmt.Printf("source=%v\n", array)
	// 22, 10, 8 60, 4
	// 60, 10, 8, 22, 4
	// 60, 22, 8, 10,4
	// 60, 22, 10, 8, 4
	for i := 0; i < length; i++ {

		for j := i + 1; j < length; j++ {

			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
				fmt.Printf("%d val=%v\n", i, array)
			}
		}
	}
}
