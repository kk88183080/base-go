package main

import (
	"bytes"
	"fmt"
	"strings"
)

func init() {
	fmt.Println("我是init方法")
}

// 定义枚举
// 每次 const 出现时，都会让 iota 初始化为0
const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	//str := 'abc' + '123'
	str := "abc" + "123"
	fmt.Println(str)

	/*str:='123'+ "abc"
	fmt.Println(str)*/

	fmt.Println(fmt.Sprintf("acb%d", 123))

	fmt.Println(strings.Join([]string{"12"}, "abc"))

	buf := bytes.Buffer{}
	buf.WriteString("123")
	buf.WriteString("abc")
	fmt.Println(buf.String())

	fmt.Println(x, y, z, k, p)
}

// nil 值。nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
func testvar() {
	//  编译不通过
	//var x = nil
	// 编译通过
	var x interface{} = nil
	fmt.Println(x)
	// 编译不通过
	//var b string = nil
	//fmt.Println(b)

	var err error = nil
	fmt.Println(err)
}
