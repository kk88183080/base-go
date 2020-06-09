package main

import "fmt"

func main() {
	_91test1()
	//_91test2()
}

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func _90f() {

}

//func _90f()  {
//
//}

// 编译出错
/*

func (m map[string]string) SetValue(key string, val string) {
	m[key] = val
}
func _91test() {
	m := map[string]string{}
	m.SetValue("l", "j")
}*/

type _91Map map[string]string

func (m _91Map) SetValue(key string, val string) {
	m[key] = val
}
func _91test1() {
	m := make(_91Map)
	m.SetValue("l", "j")
	m.SetValue("l1", "j")
	m.SetValue("l2", "j")
	m.SetValue("l3", "j")
	fmt.Println(m)
}

// 编译出错
/*func _91test2()  {
	m:= &map[string]string{}
	m.SetValue("ll", "jj")
}*/
