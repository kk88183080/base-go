package main

import (
	"./Queue"
	"fmt"
)

func main() {
	testQueue()
}

/**
测试队列
*/
func testQueue() {
	q := Queue.NewQueue()
	for i := 1; i < 100; i++ {
		q.EnQueue(i)
	}
	for i := 1; i < 100; i++ {
		fmt.Println(q.DeQueue())
	}
	fmt.Println(q.DeQueue())
}
