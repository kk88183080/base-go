package main

import (
	"fmt"
	"sync"
	"time"
)

type _601Data struct {
	sync.Mutex
}

func (d *_601Data) _601test() {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {

	var d _601Data
	var count sync.WaitGroup

	count.Add(2)

	go func() {
		defer count.Done()

		d._601test()
	}()

	go func() {
		defer count.Done()

		d._601test()
	}()

	count.Wait()
}
