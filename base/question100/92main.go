package main

import (
	"fmt"
	"math"
)

var _92x int

func init() {
	_92x++
}

func main() {
	// 编译失败。init() 函数不能被其他函数调用，包括 main() 函数
	//init()

	fmt.Println(_92x)
	_92min(1, 2)

}

func _92min(a, b int) {
	fmt.Printf("the min of %d and %d is %d \n", a, b, int64(math.Min(float64(a), float64(b))))
}

func _92min1(a, b int) {
	// 切片复制，并且返回两者长度的较小值
	min := copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("the min of %d and %d is %d \n", a, b, min)
}
