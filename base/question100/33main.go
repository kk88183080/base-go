package main

import "fmt"

func main() {
	i := 1
	s := []string{"A", "B", "C"}

	i, s[i-1] = 2, "Z"
	fmt.Println(s)
	fmt.Println(i)
}
