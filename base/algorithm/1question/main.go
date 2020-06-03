package main

import "fmt"

func main() {
	//fmt.Println(solution1(10))
	fmt.Println(solution2(2))
}

/**
解法一
*/
func solution1(i int64) bool {

	for i > 1 {
		fmt.Println(i)
		if i%2 == 0 { // 取余不为0肯定不是
			i = i / 2 // 最终会计算为1， 则会退出循环
		} else {
			return false
		}
	}

	return true
}

// 4 2 1

// 按位与进行运算之后是否为0
// 与运算 都是1才为1
func solution2(i int64) bool {
	if i < 1 {
		return false
	}
	return (i & (i - 1)) == 0
}

// 规律为
// 1 0001
// 0 0000

// 2 0010
// 1 0001

// 4 0100
// 3 0011
