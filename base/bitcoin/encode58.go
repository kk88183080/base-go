package main

import (
	"fmt"
	"math/big"
)

func main() {
	//testFor2(1, 99)
	RevserseBytes(b58Alphabet)
	fmt.Printf("%s\n", b58Alphabet)

}

var b58Alphabet = []byte("qwertyuiopasdfghjklzxcvbnm")

/**
  产生Bitcoin的钱包地址
  58 编码
*/
func Base58Encode(input []byte) []byte {

	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)

	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	return result
}

/**
58 解码
*/
func Base58Decode() {

}

/**
对数组进行反转
*/
func RevserseBytes(input []byte) {

	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}

}

/**
计算1到10之间数的和
*/
func testFor() {
	sum := 0

	for i, j := 1, 10; i < j; j, i = j-1, i+1 {
		sum += i + j
		fmt.Printf("%d + %d=%d\n", i, j, sum)
	}

	fmt.Printf("sum=%d", sum)
}

// 1 2 3 4 5
/**
计算任意区间之间数的和
*/
func testFor2(start, end int64) int64 {
	sum := int64(0)

	for i, j := start, end; i <= j; j, i = j-1, i+1 {

		if i == j {
			sum += i
			fmt.Printf("  + %d=%d\n", i, sum)
		} else {
			sum += i + j
			fmt.Printf("%d + %d=%d\n", i, j, sum)
		}
	}

	fmt.Printf("sum=%d", sum)

	return sum
}
