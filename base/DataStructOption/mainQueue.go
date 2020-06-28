package main

import (
	"./CricleQueue"
	"./Queue"
	"fmt"
)

func main() {
	//testQueue()
	//testQueueFile()
	TestCricleQueue()
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

func TestCricleQueue() {
	q := &CricleQueue.CricleQueue{}
	CricleQueue.InitCricleQueue(q)

	q1 := new(CricleQueue.CricleQueue)
	CricleQueue.InitCricleQueue(q1)

	for i := 1; i < 100; i++ {
		CricleQueue.EnQueue(q, i)
	}
	for i := 1; i < 100; i++ {
		fmt.Println(CricleQueue.DeQueue(q))
	}
	fmt.Println(CricleQueue.DeQueue(q))
}
