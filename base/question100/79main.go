package main

import "fmt"

func main() {
	//_79test()
	_79test2()
}

type _79ConfigOne struct {
	Daemon string
}

func (c *_79ConfigOne) String() string {
	return fmt.Sprintf("String print: %v", c)
}

func _79test() {
	c := &_79ConfigOne{}
	c.String()
}

func _79test2() {
	c := _79ConfigOne{}
	c.String()
}

// goroutine stack exceeds 1000000000-byte limit
// 无限递归循环，栈溢出。
// 知识点：类型的 String() 方法。如果类型定义了 String() 方法，
// 使用 Printf()、Print() 、 Println() 、 Sprintf() 等格式化输出时会自动使用 String() 方法
