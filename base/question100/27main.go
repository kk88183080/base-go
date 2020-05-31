package main

import "fmt"

func main() {
	//_27test()
	_27testMap()
}

// 枚举类的定义
type _27Direction int

const (
	North _27Direction = iota // 0
	East                      // 1
	South                     // 2
	West                      // 3
)

func (i _27Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[i]
}

func _27test() {
	fmt.Println(South)
}

type Math struct {
	x, y int
}

var s = map[string]Math{
	"foo": Math{2, 3},
}

var s1 = map[string]*Math{
	"tt": &Math{2, 3},
}

func _27testMap() {
	// 编译出错
	// 错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。
	//s["foo"].x=4
	//fmt.Println(s.["foo"].x)

	temp := s["foo"]
	temp.x = 4
	s["foo"] = temp

	fmt.Println(s["foo"].x)

	s1["tt"].x = 4
	fmt.Println(s1["tt"].x)
}
