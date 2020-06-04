package main

import "fmt"

func main() {
	//_54test()
	_54test2()
}

type N int16

func (n N) value() {
	n++
	fmt.Printf("v:%p, %v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v:%p, %v\n", n, *n)
}

func (n N) print() {
	fmt.Println(n)
}

func (n *N) printPointer() {
	fmt.Println(*n)
}

func _54test() {
	var a N = 10

	p := &a
	p1 := &p

	// 编译失败
	//p1.value()
	//p1.pointer()

	fmt.Println(p, p1)
	p.value()
	p.pointer()

}

func _54test2() {

	var n N = 20

	fmt.Println(n)

	n++
	f1 := N.print
	f1(n)

	n++
	f2 := (*N).print
	f2(&n)

	n++
	f4 := (*N).printPointer
	f4(&n)

	// 编译错误
	//n++
	//f3:=N.printPointer
	//f3(n)

	// 编译错误
	//n++
	//f3:=N.printPointer
	//f3(&n)
}

// 方法表达式。通过类型引用的方法表达式会被还原成普通函数样式，接收者是第一个参数，调用时显示传参。
// 类型可以是 T 或 *T，只要目标方法存在于该类型的方法集中就可以。
