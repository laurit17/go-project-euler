package main

import "fmt"

func PrimesUpTo(n int) []bool {
	sieve := make([]bool, n+1)

	for i := 2; i < len(sieve); i++ {
		sieve[i] = true
	}

	for i := 2; i*i < len(sieve); i++ {
		if sieve[i] {
			for j := 2; i*j < len(sieve); j++ {
				sieve[i*j] = false
			}
		}
	}

	return sieve
}

func main() {
	sieve := PrimesUpTo(2000000)
	sum := 0
	for index, value := range sieve {
		if value {
			sum += index
		}
	}
  fmt.Println(sum)
}
