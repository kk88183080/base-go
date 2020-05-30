package main

import (
	"fmt"
	"math/big"
)

func main() {

	num, b := big.NewInt(0).SetString("234", 10)
	if b {
		fmt.Println(num)
	}

	num1 := big.NewInt(0).SetBytes([]byte("234"))

	fmt.Println(num1)

	fmt.Println(big.NewInt(10).Add(num1, num))
	fmt.Println(big.NewInt(0).Add(num1, num))
	fmt.Println(big.NewInt(0).Div(num1, num))
	fmt.Println(big.NewInt(0).Mod(big.NewInt(10), big.NewInt(3)))
	fmt.Println(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(2)))
}
