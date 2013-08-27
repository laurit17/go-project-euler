package main

import (
	"fmt"
	"github.com/laurit17/go_utils/iterator_utils"
)

func sieveOfEratosthenes() func() int {
	numberStream := iterator_utils.IterStartingFrom(2)
	return func() int {
		nextNumber := numberStream()
		filter_func := func(n int) bool { return n%nextNumber != 0 }
		numberStream = iterator_utils.FilterIter(filter_func, numberStream)
		return nextNumber
	}
}

func nthPrime(n int) int {
	iter := sieveOfEratosthenes()
	for i := 0; i < n-1; i++ {
		iter()
	}
	return iter()
}

func main() {
	fmt.Println(nthPrime(10001))
}
