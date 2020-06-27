package Queue

// 队列的接口定义
type MyQueue interface {
	Size() int             //获取队列的大小
	Front() interface{}    // 第一个元素
	End() interface{}      // 最后一个元素
	IsEmpty() bool         // 是否为空
	EnQueue(v interface{}) // 加入队列
	DeQueue() interface{}  // 出队列
	Clear()                // 清空队列
}

// 队列的结构体
type Queue struct {
	dataStore []interface{} // 数据存储
	theSize   int           // 队列的大小
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.dataStore = make([]interface{}, 0)
	queue.theSize = 0
	return queue
}

func (q *Queue) Size() int {
	return q.theSize
}

// 第一个元素
func (q *Queue) Front() interface{} {

	if q.theSize == 0 {
		return nil
	}
	return q.dataStore[0]
}

// 最后一个元素
func (q *Queue) End() interface{} {
	if q.theSize == 0 {
		return nil
	}

	return q.dataStore[q.theSize-1]
}

// 是否为空
func (q *Queue) IsEmpty() bool {
	return q.theSize == 0
}

// 加入队列
func (q *Queue) EnQueue(v interface{}) {
	q.dataStore = append(q.dataStore, v)
	q.theSize++
}

// 出队列
func (q *Queue) DeQueue() interface{} {

	if q.Size() == 0 {
		return nil
	}

	data := q.dataStore[0]
	if q.Size() > 1 {
		//q.dataStore = q.dataStore[1:]
		q.dataStore = q.dataStore[1:q.Size()]
	} else { // q.Size() ==1 时的处理逻辑
		// 重新开辟内存空间
		q.dataStore = make([]interface{}, 0)
	}
	q.theSize--

	return data
}

// 清空队列
func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0, 10)
	q.theSize = 0
}
