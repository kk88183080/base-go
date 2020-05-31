package main

import "fmt"

func main() {
	//_16testsliceLengthCap()
	//_16testMap()

	//_16testWork()
	_17testWok()
}

/**
测试切片的长度及容量
*/
func _16testsliceLengthCap() {
	var a = [3]int{1, 2, 3}
	fmt.Println(len(a))
	fmt.Println(cap(a))
	b := a[:1] //i=0 长度j-i=>1-0=1 容量l-i=>3-0=3
	fmt.Printf("b length:%d, cap: %d\n", len(b), cap(b))

	c := a[:0] //i=0 j-i=0-0=0 l-i=3-0=3
	fmt.Printf("c length:%d, cap: %d\n", len(c), cap(c))

	d := a[:2] //i=0 j-i=2-0=2 l-i=3-0=3
	fmt.Printf("d length:%d, cap: %d\n", len(d), cap(d))

	e := a[1:2:cap(a)] // i=1 j-i=2-1=1 3-1=2
	fmt.Printf("e length:%d, cap: %d\n", len(e), cap(e))
}

/*
知识点：数组或切片的截取操作。截取操作有带 2 个或者 3 个参数，
形如：[i:j] 和 [i:j:k]，
假设截取对象的底层数组长度为 l。
在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
*/

func _16testMap() {
	// 只是定义，没有给初始化
	//var m map[string]int

	var m = make(map[string]int)
	// 当m没有初始化进，会报错 :panic: assignment to entry in nil map
	m["a"] = 1
	// 错误的写法，m获取时，回返回两个参数，第一个是值，第二个是bool; 值存在时会返回true, 否则返回false
	//if v := m["b"]; v != nil {
	//	fmt.Printf(v)
	//}

	if v, ok := m["b"]; ok {
		fmt.Printf("b is exist, v=%d\n", v)
	} else {
		fmt.Printf("b is no exist\n")
	}

	if v, ok := m["a"]; ok {
		fmt.Printf("a is exists v=%d\n", v)
	} else {
		fmt.Printf("a is not exists\n")
	}
}

//
type _16A interface {
	ShowA() int
}

type _16B interface {
	ShowB() int
}

type _16work struct {
	i int
}

func (w _16work) ShowA() int {
	return w.i + 10
}

func (w _16work) ShowB() int {
	return w.i + 20
}

func _16testWork() {
	//w:=_16work{3}
	//
	//var a _16A = w
	//var b _16B = w

	// 知识点：接口的静态类型。a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；
	// a 的静态类型是 A，b 的静态类型是 B，
	// 接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错。看下编译错误：
	//fmt.Println(a.ShowB())
	//fmt.Println(b.ShowA())
}

func _17testWok() {
	var a _16A = _16work{3}

	c := a.(_16work)

	fmt.Println(c.ShowA())
	fmt.Println(c.ShowB())
}
