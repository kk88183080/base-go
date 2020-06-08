package main

import "fmt"

func main() {
	//_86test()
	//_86test1()
	_86test2()
}

func _86test() {
	n := 43210
	fmt.Println(n/60*60, "hours and ", n%60*60, "seconds")

}

func _86test1() {
	n := 43210
	fmt.Println(n/(60*60), "hours and ", n%(60*60), "seconds")

}

const (
	C = 100
	D = 010 // 8
	F = 001 // 1
)

func _86test2() {
	fmt.Println(D)
	fmt.Println(F)
	fmt.Println(C + 2*D + 2*F) // 100 + 16 + 2
}
