package main

import (
	"fmt"
	"os"
)

func main() {
	//TestPrints()
	//TestFprints()
	TestSprints()
}

/**
输出到标准输出
*/
func TestPrints() {
	// 输出的没空格, 有数字时才有空格
	fmt.Print("xxx", "zzz", 1, 2, 3)
	fmt.Print("\n")

	// 输出的有空格
	fmt.Println("aaa", "bbb", 2, 3, "ccc")

	//
	fmt.Printf("%s, %s, %d, %d, %s\n", "eee", "fff", 1, 2, "ggg")
}

/**
设置输出到指定的输出上
*/
func TestFprints() {
	// 标准错误输出的级别标准输出的高
	// 打印的输出没有空格,不换行；当有数字时才有空格
	fmt.Fprint(os.Stdout, "a", "b", "c", 1, 2, 3, "\n")
	fmt.Fprint(os.Stderr, "Stderr", "ddd", "ccc", "\n")

	// 把数据输出到标准输入，所以界面上看不到任何显示
	fmt.Fprint(os.Stdin, "zzz", "\n")

	// 打印的输出有空格，也有换行
	fmt.Fprintln(os.Stdout, "a", "b", "c", 1, 2, 3)
	fmt.Fprintln(os.Stderr, "a", "b", "c", 4, 5, 6)

	// 占位符的数量和输入的参数
	fmt.Fprintf(os.Stderr, "%s %d %d %d %s %s", "f", 1, 2, 3, "g", "i")
}

/**
返回string对象
*/
func TestSprints() {
	rs := fmt.Sprint("a", "b", "c")
	fmt.Printf("%s\n", rs)

	rs = fmt.Sprint("a", "b", "c", 2, 3, 5)
	fmt.Printf("%s\n", rs)

	rs = fmt.Sprintln("a", "1", "c")
	fmt.Printf("%s", rs)

	rs = fmt.Sprintf("%s 你好！", "ljw")
	fmt.Printf("%s", rs)
}
