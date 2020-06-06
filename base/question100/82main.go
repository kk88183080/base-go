package main

import "fmt"

func main() {
	//_82test()
	_82test2()
}

const (
	azero = iota //0
	aone  = iota //1
)

const (
	info  = "msg"
	bzero = iota // 1
	bzone = iota // 2
)

func _82test() {
	fmt.Println(azero, aone)
	fmt.Println(bzero, bzone)
}

func _82test2() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)

		if n == -n {
			fmt.Println("n==-n", n, -n)
			count++
		}

		if m == -m {
			fmt.Println("m==-m", m, -m)
			count++
		}
	}

	fmt.Println(count)
}

//数值溢出。当 i 的值为 0、128 是会发生相等情况，注意 byte 是 uint8 的别名
