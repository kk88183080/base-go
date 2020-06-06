package main

import "fmt"

func main() {
	_83test()
}

type _83data struct {
	name string
}

func (d *_83data) print() {
	fmt.Println("name:", d.name)
}

type _83potiner interface {
	print()
}

func _83test() {
	d1 := _83data{"one"}
	d1.print()

	// 只能给接口赋值指针数据
	var d2 _83potiner = &_83data{"two"}
	d2.print()
}
