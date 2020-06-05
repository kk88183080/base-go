package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//d := _602{new(sync.Mutex)}
	var d _602
	d.Mutex = new(sync.Mutex)
	c := sync.WaitGroup{}
	c.Add(2)

	go func() {
		defer c.Done()
		d._602test()
	}()

	go func() {
		defer c.Done()
		d._602test()
	}()

	c.Wait()
}

type _602 struct {
	*sync.Mutex
}

func (d _602) _602test() {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
