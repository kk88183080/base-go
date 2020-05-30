package main

import "fmt"

/*func main() {

	// new 之后的对象是一个指针类型的值
	l := new([]int)
	// append 不能对指针类型的参数做操作， 要想编译通过；可以使用make关键字
	l = append(l, 1)
	fmt.Println(l)
}*/

/*func main() {
	s1:=[]int{1, 2, 3}
	s2:= []int{4,5}

	// 要把第二个切片的值 ，拼装到第一个切片的后面；就是加...;要不编译不通过
	s3:=append(s1, s2...)

	fmt.Println(s3)

}*/

var (
	// 这块用:=会出错，编译不能通过；使用var就不能用:=
	// := 只能在函数内部使用
	//size := 1024
	size     = 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}
