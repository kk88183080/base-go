package main

import "fmt"

func main() {
	test12()
}

// %d表示输出十进制数字，+表示输出数值的符号。这里不表示取反。
func test12() {
	i := -5
	j := +5
	fmt.Printf("%+d %+d", i, j)
	fmt.Println()
	fmt.Printf("%d %d", i, j)
}
