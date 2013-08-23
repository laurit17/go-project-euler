package main

import "fmt"

func sumOfEvenFibsUpTo(n int) int {
	sum := 0
	a, b := 1, 0
	for i := 0; ; i++ {
		a, b = a+b, a
		if b >= n {
			break
		}
		if b%2 == 0 {
			sum += b
		}
	}

	return sum
}

func main() {
	const MAX = 4000000
	fmt.Printf("%d\n", sumOfEvenFibsUpTo(MAX))
}
