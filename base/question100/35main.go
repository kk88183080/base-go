package main

import "fmt"

func main() {
	//_35testbool()
	//_35iAddAdd()

	setTest()
}

func _35testbool() {
	var b bool
	b = true
	fmt.Println(b)

	// 编译出错
	//b=1

	// 编译出错
	//b = bool(1)
	b = false
	fmt.Println(b)

	b = (1 == 2)
	fmt.Println(b) // false
}

func _35iAddAdd() {

	i := 1
	i++
	fmt.Println(i)
	// 编译出错
	//j=i++

	// 编译出错
	//++i

	// 编译出错
	//--i

	i--
	fmt.Println(i)
}

//自增自减操作。i++ 和 i-- 在 Go 语言中是语句，不是表达式，因此不能赋值给另外的变量。此外没有 ++i 和 --i。

type Trasinfo struct {
}

type Fragment interface {
	Exec(transinfo *Trasinfo) error
}

type GetPodAction struct {
}

func (g GetPodAction) Exec(t *Trasinfo) error {
	return nil
}

func setTest() {
	// 编译错误
	var f1 Fragment = new(GetPodAction)

	// 编译错误
	//var f Fragment = GetPodAction

	// 编译错误
	var f2 Fragment = &GetPodAction{}

	//
	var f3 Fragment = GetPodAction{}

	fmt.Println(f1, f2, f3)
}
