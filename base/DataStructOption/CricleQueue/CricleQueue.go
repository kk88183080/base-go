package CricleQueue

import "errors"

const SIZE = 100 // 最多存储SIZE-1个元素

type CricleQueue struct {
	dataStore [SIZE]interface{} // 队列的大小
	front     int               // 头指针
	rear      int               // 尾指针
}

func InitCricleQueue(q *CricleQueue) {
	q.front = 0
	q.rear = 0
}

func QueueLength(q *CricleQueue) int {

	return (q.rear - q.front + SIZE) % SIZE
}

func EnQueue(q *CricleQueue, v interface{}) error {

	if (q.rear+1)%SIZE == (q.front)%SIZE {
		return errors.New("队列已满")
	}

	q.dataStore[q.rear] = v
	q.rear = (q.rear + 1) % SIZE

	return nil
}

func DeQueue(q *CricleQueue) (interface{}, error) {

	if q.front == q.rear {
		return nil, errors.New("队列为空")
	}

	res := q.dataStore[q.front] // 获取数据
	q.dataStore[q.front] = 0    //
	q.front = (q.front + 1) % SIZE

	return res, nil
}
