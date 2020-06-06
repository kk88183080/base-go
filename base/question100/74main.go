package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	//_74test()
	_74test2()
}

func _74test() {

	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	for {
	}
}

// 虽然提示错误，便是可以执行
// for {} 独占 CPU 资源导致其他 Goroutine 饿死。

func _74test2() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()

	select {}
}
