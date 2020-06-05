package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//_60test()
	_60test3()
}

type _60T struct {
}

func (t *_60T) _60foo() {

}

func (_60T) _60bar() {

}

type _60S struct {
	*_60T
}

func _60test() {
	s := _60S{}

	_ = s._60foo //
	s._60foo()

	_ = s._60bar //这块会出错 panic(1)
}

//因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic。

type _60data struct {
	sync.Mutex
}

func (n _60data) _60test2(s string) {
	n.Lock()
	defer n.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func _60test3() {

	var wg sync.WaitGroup
	wg.Add(2)

	var data _60data

	go func() {
		defer wg.Done()

		data._60test2("read")
	}()

	go func() {

		defer wg.Done()
		data._60test2("write")
	}()

	wg.Wait()
}

// 锁失效。将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。
