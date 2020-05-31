package main

import "fmt"

func main() {
	//_26test()

	_26testInter()
}

const (
	a = iota // 0
	b = iota // 1
)

const (
	name = "name" // 0
	c    = iota   // 1
	d    = iota   // 2
)

func _26test() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// 编译出错
	//fmt.Println(iota)
}

//iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。
//iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。

type _26people interface {
	show()
}

type _26student struct {
}

func (stu *_26student) show() {

}

func _26testInter() {
	var s *_26student

	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}

	var p _26people = s

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}

//这道题会不会有点诧异，我们分配给变量 p 的值明明是 nil，然而 p 却不是 nil。
//记住一点，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
//上面的代码，给变量 p 赋值之后，p 的动态值是 nil，但是动态类型却是 *Student，是一个 nil 指针，所以相等条件不成立
