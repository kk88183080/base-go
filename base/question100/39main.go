package main

import (
	"fmt"
	"time"
)

func main() {
	//_39test2()
	//_39chantest()

	_39chantest2()
}

func _39test(x interface{}) {

	if x == nil {
		fmt.Println("empty interface")
		return
	}

	fmt.Println("non empty interface")
}

func _39test2() {

	var x *int = nil
	_39test(x)
}

//考点：interface 的内部结构，我们知道接口除了有静态类型，还有动态类型和动态值，
//当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。这里的 x 的动态类型是 *int，所以 x 不为 nil

func _39chantest() {
	ch := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			i, ok := <-ch

			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println(i)
		}
	}()

	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 3)
}

//程序抛异常。先定义下，第一个协程为 A 协程，第二个协程为 B 协程；当 A 协程还没起时，主协程已经将 channel 关闭了，
//当 A 协程往关闭的 channel 发送数据时会 panic，panic: send on closed channel。

// 可以正确执行的代码
func _39chantest2() {
	ch := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			i, ok := <-ch

			if !ok {
				fmt.Println("close")
				close(ch)
				return
			}
			fmt.Println(i)
		}
	}()

	fmt.Println("ok")
	time.Sleep(time.Second * 3)
}
