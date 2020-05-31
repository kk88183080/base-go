package main

import (
	"fmt"
	"time"
)

func main() {
	//_29testslice()

	//_29testslice1()
	//_29testslice2()
	_29testslice3()
}

func _29testslice() {
	v := []int{1, 2, 3}
	for i, val := range v {
		fmt.Println(i, val)
		v = append(v, i)
	}

	fmt.Println(v)
}

func _29testslice1() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

func _29testslice2() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func(i int, v int) {
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(time.Second * 3)
}

func _29testslice3() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		it := i
		vt := v
		go func() {
			fmt.Println(it, vt)
		}()
	}

	time.Sleep(time.Second * 3)
}
