package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func fibonacciBig(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}
	prev := big.NewInt(0)
	curr := big.NewInt(1)
	for i := 2; i <= n; i++ {
		prev, curr = curr, new(big.Int).Add(prev, curr)
	}
	return curr
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Please provide a valid number as an argument")
	}

	start := time.Now()
	res := fibonacciBig(num)
	duration := time.Since(start)
	fmt.Println(res)
	fmt.Printf("Time taken: %v", duration)
}
