package main

import (
	"fmt"
	"github.com/laurit17/go_utils/math_utils"
)

func main() {
	sieve := math_utils.PrimesUpTo(2000000)
	sum := 0
	for index, value := range sieve {
		if value {
			sum += index
		}
	}
	fmt.Println(sum)
}
