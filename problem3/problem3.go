package main

import (
	"fmt"
	"math"
)

func isPrimeInt64(num int64) bool {
	if num < 2 {
		return false
	}
	for i := int64(2); i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findBiggestPrimeFactor(big int64) int64 {
	root := int64(math.Sqrt(float64(big)))

	if root%2 == 0 {
		root -= 1
	}
	for i := root; i > 0; i -= 2 {
		if big%i == 0 && isPrimeInt64(i) {
			return i
		}
	}
	return 1
}

func main() {
	const big = int64(600851475143)
	fmt.Printf("%d\n", findBiggestPrimeFactor(big))
}
