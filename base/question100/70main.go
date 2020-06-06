package main

import "fmt"

func main() {
	_70test()
}

//引号、双引号和字符串连接。在 Go 语言中，双引号用来表示字符串 string，其实质是一个 byte 类型的数组，单引号表示 rune 类型。

func _70defer1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()

	return r
}

func _70defer2(i int) (r int) {
	defer func() {
		r += i
	}()

	return 2
}

func _70test() {
	fmt.Println(_70defer1(1)) // 4
	fmt.Println(_70defer2(1)) // 3
}
