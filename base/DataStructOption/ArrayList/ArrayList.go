package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 获取指定下标的元素
	Set(index int, val interface{}) error       // 更新指定下标中的元素
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newVal interface{})                  // 追加数据
	Clear() error                               // 清空数组
	Delete(index int) error                     // 删除第几个元素
	String() string                             // 返回字符串
	Iterator() Iterator                         //实现迭代器
}

// 数据结构，字符串，数字，实数
type ArrayList struct {
	dataStore []interface{} // 存储数据的数组
	theSize   int           // 数组的长度
}

// 开辟内存
func NewArrayList() *ArrayList {

	a := new(ArrayList) // 初始化结构体

	a.dataStore = make([]interface{}, 0, 10) // 开辟10个空间
	a.theSize = 0

	return a
}

/**
获取数组长度
*/
func (a *ArrayList) Size() int {

	return a.theSize
}

func (a *ArrayList) Get(index int) (interface{}, error) {

	if a.theSize <= index || index < 0 {
		return nil, errors.New("索引越界")
	}

	return a.dataStore[index], nil
}

// 更新指定下标中的元素
func (a *ArrayList) Set(index int, val interface{}) error {

	if index >= a.theSize || index < 0 {
		return errors.New("索引越界")
	}

	a.dataStore[index] = val
	return nil
}

// 插入数据
func (a *ArrayList) Insert(index int, newVal interface{}) error {

	if index >= a.theSize || index < 0 {
		return errors.New("索引越界")
	}

	// 监测内存是否要扩容
	a.checkIsFull()
	a.dataStore = a.dataStore[:a.theSize+1] // 内存向后移动一位

	for i := a.theSize; i > index; i-- {
		a.dataStore[i] = a.dataStore[i-1]
	}

	a.dataStore[index] = newVal
	a.theSize++

	return nil
}

func (a *ArrayList) checkIsFull() { // 扩容

	if a.theSize == cap(a.dataStore) {
		// 一定要设置长度，不然没法复制数据
		newDataStore := make([]interface{}, a.theSize*2, a.theSize*2) // 扩容成原来的2倍

		// 方法1
		copy(newDataStore, a.dataStore)

		// 方法2
		//for i := 0; i < a.theSize; i++ {
		//	newDataStore[i] = a.dataStore[i]
		//}

		a.dataStore = newDataStore
	}
}

func (a *ArrayList) Append(newVal interface{}) {

	a.dataStore = append(a.dataStore, newVal)
	a.theSize++

}

// 清空数组
func (a *ArrayList) Clear() error {

	a.dataStore = make([]interface{}, 0, 10)
	a.theSize = 0

	return nil
}

// 删除第几个元素
func (a *ArrayList) Delete(index int) error {

	a.dataStore = append(a.dataStore[0:index], a.dataStore[index+1:]...)
	a.theSize--

	return nil
}

func (a *ArrayList) String() string {
	return fmt.Sprint(a.dataStore)
}
