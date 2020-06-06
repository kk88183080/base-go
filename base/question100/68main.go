package main

import "fmt"

func main() {
	_68test1()
}

/*func _68test() {

	for i := 0; i < 10; i++ {
	loop:
		fmt.Println(i)
	}

	goto loop
}*/

//goto 不能跳转到其他函数或者内层代码。编译报错

func _68test1() {
	x := []int{0, 1, 2}
	y := [3]*int{}

	for i, v := range x {
		defer func() {
			fmt.Println(v) // 2 2 2
		}()

		y[i] = &v
	}

	fmt.Println(*y[0], *y[1], *y[2]) // 2 2 2
}

//defer()、for-range。for-range 虽然使用的是 :=，但是 v 不会重新声明，可以打印 v 的地址验证下。
