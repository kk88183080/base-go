package ArrayList

import "errors"

type Iterator interface {
	HashNext() bool // 是否有下个

	Next() (interface{}, error) // 获取下一个元素

	Remove() // 移除

	GetIndex() int // 获取当前索引
}

type Iterable interface {
	Iterator() Iterator // 初始化迭代器
}

type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前的数组索引
}

func (a *ArrayList) Iterator() Iterator {

	it := new(ArrayListIterator)
	it.list = a
	it.currentIndex = 0

	// Iterator 一定要实现它的所有接口
	return it
}

// 是否有下个
func (i *ArrayListIterator) HashNext() bool {
	return i.currentIndex <= i.list.theSize-1
}

// 获取下一个元素
func (i *ArrayListIterator) Next() (interface{}, error) {

	if i.HashNext() {
		val, _ := i.list.Get(i.currentIndex)
		i.currentIndex++
		return val, nil
	}

	return nil, errors.New("没有下一个了")
}

// 移除
func (i *ArrayListIterator) Remove() {

	if i.currentIndex > 0 {
		i.currentIndex--
	}
	i.list.Delete(i.currentIndex)
}

// 获取当前索引
func (i *ArrayListIterator) GetIndex() int {
	return i.currentIndex
}
