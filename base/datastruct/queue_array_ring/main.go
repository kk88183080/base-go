package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	q := Queue{
		4,
		0,
		0,
		[4]int{},
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入push添加元素")
		fmt.Println("2.输入pop弹出元素")
		fmt.Println("3.输入show查看元素")
		fmt.Println("4.输入exit退出")

		fmt.Scanln(&key)
		switch key {
		case "push":
			fmt.Println("请输入要添加的元素")

			fmt.Scanln(&val)

			e := q.push(val)
			if e != nil {
				fmt.Println("push出错:", e.Error())
			} else {
				fmt.Println("push成功")
			}
		case "pop":
			v, e := q.pop()
			if e != nil {
				fmt.Println("pop出错:", e.Error())
			} else {
				fmt.Println("pop值为:", v)
			}
		case "show":
			e := q.listShow()
			if e != nil {
				fmt.Println("查看出错：", e.Error())
			}
		case "exit":
			os.Exit(0)
		}

	}
}

type Queue struct {
	maxsize int
	head    int // 默认值为0
	tail    int // 默认值为0
	data    [4]int
}

func (q *Queue) push(v int) (err error) {
	if q.isFull() {
		fmt.Println("队列已满")
		return errors.New("队列已满")
	}

	q.data[q.tail%q.maxsize] = v
	q.tail++
	return
}

func (q *Queue) pop() (v int, err error) {
	if q.isEmpty() {
		fmt.Println("队列为空，不能获取")
		return -1, errors.New("队列为空，不能获取")
	}

	v = q.data[q.head%q.maxsize]
	q.head++
	return v, nil
}

func (queue *Queue) listShow() (err error) {
	if queue.isEmpty() {
		return errors.New("队列为空")
	}

	fmt.Printf("开始打印,head:%d, tail:%d\n", queue.head, queue.tail)

	start := queue.head % queue.maxsize
	for i := start; i < (queue.maxsize - 1); i++ {
		fmt.Print(queue.data[i])
	}
	fmt.Println()

	return
}

func (queue *Queue) isFull() bool {
	return ((queue.tail + 1) % queue.maxsize) == (queue.head % queue.maxsize)
}

func (queue *Queue) getSize() int {
	return (queue.tail + queue.maxsize - queue.head) % queue.maxsize
}

func (queue *Queue) isEmpty() bool {
	return queue.head == queue.tail
}
