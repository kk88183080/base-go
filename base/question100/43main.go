package main

/*import (
	"fmt"
	"log"
	"time"
)

func main() {

}*/

//入的包没有被使用。如果引入一个包，但是未使用其中如何函数、接口、结构体或变量的话，代码将编译失败。
//如果你真的需要引入包，可以使用下划线操作符，_，来作为这个包的名字，从而避免失败。下划线操作符用于引入，但不使用。
//我们还可以注释或者移除未使用的包。

import (
	"fmt"
	_ "fmt"
	"log"
	"time"
)

func main() {
	//_43test()
	//_43test2()
	_43test3()
}

var _ = log.Println

func _43test() {
	_ = time.Now()
}

// 如果要强制使用，可以使用_来处理，但会执行引入包中的init方法
// main函数一定要定在import下面，不能在前面；否则会编译出错

func _43test2() {
	x := interface{}(nil)
	y := (*int)(nil)

	a := y == x
	b := y == nil

	_, c := x.(interface{})
	fmt.Println(a, b, c)
}

// 类型断言。类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。
// 编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败

func _43test3() {
	var s []int
	// 切片使用make时，必须传入长度参数
	//s=make([]int,0)
	s = append(s, 1)

	var m map[string]int
	//m= make(map[string]int)
	m["x"] = 1

	fmt.Println(s)
	fmt.Println(m)
}

//有 1 出错误，不能对 nil 的 map 直接赋值，需要使用 make() 初始化。但可以使用 append() 函数对为 nil 的 slice 增加元素。
