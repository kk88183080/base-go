package main

func main() {

}

type _48foo struct {
	bar int
}

/*func _48test() {
	var b _48foo

	b.bar, t:= 1, 10

	fmt.Println(b.bar, t)
}*/
//:= 操作符不能用于结构体字段赋值

func _48test1() {
	//fmt.Println(~2)
}

// 编译错误
//很多语言都是采用 ~ 作为按位取反运算符，Go 里面采用的是 ^ 。按位取反之后返回一个每个 bit 位都取反的数，
//对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），对于无符号整数来说就是按位取反
