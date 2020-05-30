package main

import "fmt"

func main() {
	var m map[_9person]int
	p := _9person{name: "ljw"}
	fmt.Println(m[p])

	var m1 map[_9person]string
	fmt.Println(m1[p])

	var m2 map[_9person]bool
	fmt.Println(m2[p])

	var m3 map[_9person]float64
	fmt.Println(m3[p])
}

type _9person struct {
	name string
}
