package main

import (
	"fmt"
	"sync"
)

func main() {
	//_76test1()
	_76test2()
}

func _76test() {
	var wg sync.WaitGroup

	wg.Add(1)
	wg.Add(1)
	go func() {
		wg.Done()
	}()

	wg.Wait()
}

// 这个会死锁

func _76test1() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		wg.Add(1)
	}()

	wg.Wait()
}

// WaitGroup is reused before previous Wait has returned

type _76S struct {
}

func (s _76S) f() {
	fmt.Println("s.f")
}

func (s _76S) g() {
	fmt.Println("s.g")
}

type _76S2 struct {
	_76S
}

func (s _76S2) f() {
	fmt.Println("s2.f")
}

type _76I interface {
	f()
}

func _76pringtype(i _76I) {
	fmt.Printf("%T\n", i)

	if s1, ok := i.(_76S); ok {
		s1.f()
		s1.g()
	}

	if s2, ok := i.(_76S2); ok {
		s2.f()
		s2.g()
	}

}

func _76test2() {
	_76pringtype(_76S{})
	_76pringtype(_76S2{})
}
