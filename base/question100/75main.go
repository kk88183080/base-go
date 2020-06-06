package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//_75test()
	_75test2()
}

func _75test() {
	f, err := os.Open("file")

	if err != nil {
		return
	}

	defer f.Close()

	val, err := ioutil.ReadAll(f)

	if err != nil {
		return
	}

	fmt.Println(string(val))
}

func _75test2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v, %T", r, r)
		}
	}()
	panic(1)
	panic(2)
}

//panic、recover()。当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容。
