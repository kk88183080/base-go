package main

import (
	"fmt"
	"sync"
)

func main() {
	//_80test()
	//_80test1()
	//_80test2()
	//_80test3()
	_80test4()
}

func _80test() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()
}

func _80test1() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()
}

// panic: sync: negative WaitGroup counter
func _80test2() {

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()
}

// panic: sync: negative WaitGroup counter

func _80test3() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func _80test4() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}

	wg.Wait()
}
