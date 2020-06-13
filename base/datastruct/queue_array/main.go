package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	q := Queue{4, [4]int{}, -1, -1}

	var key string
	var v int
	for {
		fmt.Println("1. 输入add代表添加数据到队列")
		fmt.Println("2. 输入get代表从队列中获取数据")
		fmt.Println("3. 输入show代表查看队列中的数据")
		fmt.Println("4. 输入exit代表退出当前程序")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入要添加到队列中的值")
			fmt.Scanln(&v)
			e := q.Add(v)
			if e != nil {
				fmt.Println("队列添加失败", e.Error())
			} else {
				fmt.Println("队列添加成功")
			}
		case "get":
			v, e := q.get()
			if e != nil {
				fmt.Println("队列获取失败", e.Error())
			} else {
				fmt.Printf("队列获取成功，v=%d\n", v)
			}
		case "show":
			q.show()
		case "exit":
			os.Exit(0)
		}
	}
}

/**
 定义队列的数据结构
使用数组来处理
*/
type Queue struct {
	maxSize int    // 队列的最大容量
	array   [4]int // 使用数组进行数据存储
	front   int    // 队列的头部，默认为-1
	tail    int    // 队列的尾部，默认为-1
}

/**
添加时front值不变化，tail每次加1
*/
func (q *Queue) Add(v int) (err error) {

	if q.tail+1 >= q.maxSize {
		fmt.Println("队列已满")
		return errors.New("队列已满")
	}

	q.tail++
	q.array[q.tail] = v
	return
}

func (q *Queue) show() (err error) {

	if q.tail == -1 {
		fmt.Println("队列为空")
		return errors.New("队列为空")
	}

	for i := q.front + 1; i <= q.tail; i++ {
		fmt.Print(q.array[i], " ")
	}
	fmt.Println(" 打印完成")
	return
}

func (q *Queue) get() (v int, err error) {
	if q.tail == q.front {
		return -1, errors.New("队列为空")
	}

	q.front++

	return q.array[q.front], nil
}
