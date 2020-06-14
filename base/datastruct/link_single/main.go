package main

import "fmt"

func main() {
	h := HeroNode{}

	node1 := HeroNode{1, "l", nil}
	h.add(node1)

	node2 := HeroNode{2, "j", nil}
	h.add(node2)

	node3 := HeroNode{3, "w", nil}
	h.add(node3)

	node4 := HeroNode{4, "z", nil}
	h.add(node4)

	node5 := HeroNode{5, "a", nil}
	h.add(node5)

	fmt.Println("add len:", h.getSize())
	fmt.Println("add data:", h.getList())

	h.del(node1)
	h.del(node2)
	h.del(node3)
	h.del(node4)
	h.del(node5)

	fmt.Println("del len:", h.getSize())
	fmt.Println("del data:", h.getList())

	h.add(node1)
	h.add(node2)
	h.add(node3)
	h.add(node4)
	h.add(node5)
	fmt.Println("add2 len:", h.getSize())
	fmt.Println("add2 data:", h.getList())

	h.del(node1)
	h.del(node2)
	h.del(node3)
	h.del(node4)
	h.del(node5)

	fmt.Println("del2 len:", h.getSize())
	fmt.Println("del2 data:", h.getList())
}

type HeroNode struct {
	no   int
	name string
	next *HeroNode
}

/**
最后添加元素
*/
func (h *HeroNode) add(node HeroNode) (e error) {

	temp := h

	for {
		if temp.next == nil { // 尾部没有元素时，就认为是最后一个节点
			temp.next = &node
			break
		} else {
			temp = temp.next
		}
	}

	return
}

/**
删除no号相同的节点
*/
func (h *HeroNode) del(node HeroNode) (ok bool) {
	ok = false

	up := h
	temp := h.next
	for {
		if temp == nil {
			break
		} else if temp.no == node.no {
			up.next = temp.next
			temp.next = nil // 设置为空，让gc去回收
			ok = true
			break
		} else {
			up = temp
			temp = temp.next
		}
	}

	return ok
}

/**
获取元素个数
*/
func (h *HeroNode) getSize() (s int) {
	i := 0
	temp := h.next

	for {
		if temp == nil {
			break
		} else {
			temp = temp.next
			i++
		}
	}
	return i
}

/**
获取元素列表
*/
func (h *HeroNode) getList() []HeroNode {
	nodes := []HeroNode{}

	temp := h.next
	for {
		if temp == nil {
			break
		} else {
			nodes = append(nodes, *temp)
			temp = temp.next
		}
	}

	return nodes
}
