package Link

type Node struct {
	data  interface{}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(v interface{})
	Pop() interface{}
	Length() int
}

/**
实例化方法
*/
func NewListStack() *Node {
	return &Node{}
}

func (l *Node) IsEmpty() bool {
	return l.pNext == nil && l.data == nil
}

/**
尾部插入
*/
func (l *Node) PushLast(v interface{}) {

	// 第一次添加数据
	if l.pNext == nil && l.data == nil {
		l.data = v
		return
	}

	temp := l.pNext
	for {
		if temp == nil {
			temp.pNext = NewListStack()
			temp.pNext.data = v
		} else {
			temp = temp.pNext
		}
	}

}

/**
队首添加节点
*/
func (l *Node) Push(v interface{}) {
	// 第一次插入数据
	if l.data == nil && l.pNext == nil {
		l.data = v
		return
	}

	temp := l.pNext

	if temp == nil {
		l.pNext = NewListStack()
		l.pNext.data = v
	} else {
		l.pNext = NewListStack()
		l.pNext.data = v

		l.pNext.pNext = temp
	}
}

/**
队首弹出节点
*/
func (l *Node) Pop() {

}
