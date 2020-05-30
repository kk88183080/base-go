package main

import "fmt"

func main() {

	h := hello
	//h := hello()
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func hello() []string {
	return nil
}
