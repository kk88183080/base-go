package main

import (
	"./Queue"
	"fmt"
)

func main() {
	//testQueue()
	testQueueFile()
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

/**
查询文件
*/
func testQueueFile() {
	Queue.QueryFile("/Users/liangweili/Desktop/bigdata-workspace/base-go/base/DataStructOption")
}
