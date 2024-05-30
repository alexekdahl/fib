package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
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

func recursiveFibonacciWithCache(n int, cache map[int]*big.Int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}
	if cache[n] != nil {
		return cache[n]
	}
	cache[n] = new(big.Int).Add(recursiveFibonacciWithCache(n-1, cache), recursiveFibonacciWithCache(n-2, cache))
	return cache[n]
}

func recursiveFibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}

func multiplyMatrix(a, b [2][2]*big.Int) [2][2]*big.Int {
	return [2][2]*big.Int{
		{
			new(big.Int).Add(new(big.Int).Mul(a[0][0], b[0][0]), new(big.Int).Mul(a[0][1], b[1][0])),
			new(big.Int).Add(new(big.Int).Mul(a[0][0], b[0][1]), new(big.Int).Mul(a[0][1], b[1][1])),
		},
		{
			new(big.Int).Add(new(big.Int).Mul(a[1][0], b[0][0]), new(big.Int).Mul(a[1][1], b[1][0])),
			new(big.Int).Add(new(big.Int).Mul(a[1][0], b[0][1]), new(big.Int).Mul(a[1][1], b[1][1])),
		},
	}
}

func matrixPower(matrix [2][2]*big.Int, n int) [2][2]*big.Int {
	result := [2][2]*big.Int{
		{big.NewInt(1), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1)},
	}
	base := matrix

	for n > 0 {
		if n&1 == 1 {
			result = multiplyMatrix(result, base)
		}
		base = multiplyMatrix(base, base)
		n >>= 1
	}

	return result
}

func bitManipulationFibonacci(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	matrix := [2][2]*big.Int{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}

	result := matrixPower(matrix, n-1)

	return result[0][0]
}

func fibonacciBitManipulationWithCache(n int, cache map[int]*big.Int) *big.Int {
	if val, ok := cache[n]; ok {
		return val
	}

	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	matrix := [2][2]*big.Int{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}

	result := matrixPower(matrix, n-1)
	fib := result[0][0]

	cache[n] = fib

	return fib
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Please provide a valid number as an argument")
	}

	var wg sync.WaitGroup
	wg.Add(5)
	results := make(chan string, 5)

	// Iterative Fibonacci
	go func() {
		defer wg.Done()
		timeStart := time.Now()
		_ = fibonacciBig(num)
		duration := time.Since(timeStart)
		results <- fmt.Sprintf("Iterative Fibonacci:\nBig O: O(n)\nfibonacci(%d)\nTime taken to calculate: %v\n", num, duration)
	}()

	// Recursive Fibonacci with cache
	recCache := make(map[int]*big.Int)
	go func() {
		defer wg.Done()
		timeStart := time.Now()
		_ = recursiveFibonacciWithCache(num, recCache)
		duration := time.Since(timeStart)
		results <- fmt.Sprintf("Recursive Fibonacci with cache:\nBig O: O(n)\nfibonacci(%d)\nTime taken to calculate: %v\n", num, duration)
	}()

	// Recursive Fibonacci without cache
	go func() {
		defer wg.Done()
		timeStart := time.Now()
		_ = recursiveFibonacci(num)
		duration := time.Since(timeStart)
		results <- fmt.Sprintf("Recursive Fibonacci without cache:\nBig O: O(2^n)\nfibonacci(%d)\nTime taken to calculate: %v\n", num, duration)
	}()

	// Bit Manipulation Fibonacci
	go func() {
		defer wg.Done()
		timeStart := time.Now()
		_ = bitManipulationFibonacci(num)
		duration := time.Since(timeStart)
		results <- fmt.Sprintf("Bit Manipulation Fibonacci:\nBig O: O(log n)\nfibonacci(%d)\nTime taken to calculate: %v\n", num, duration)
	}()

	// Bit Manipulation Fibonacci with bitCache
	bitCache := make(map[int]*big.Int)
	go func() {
		defer wg.Done()
		timeStart := time.Now()
		_ = fibonacciBitManipulationWithCache(num, bitCache)
		duration := time.Since(timeStart)
		results <- fmt.Sprintf("Bit Manipulation Fibonacci with cache:\nBig O: O(log n)\nfibonacci(%d)\nTime taken to calculate: %v\n", num, duration)
	}()

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}
}
