package main

import "fmt"

func triangleSum(low int, high int) int {
	return high*(high+1)/2 - low*(low+1)/2
}

func divisibleSumLessThan(max int, divisor int) int {
	return triangleSum(0, (max-1)/divisor) * divisor
}

func main() {
	threes_sum := divisibleSumLessThan(1000, 3)
	fives_sum := divisibleSumLessThan(1000, 5)
	fifteens_sum := divisibleSumLessThan(1000, 15)
	sum := threes_sum + fives_sum - fifteens_sum
	fmt.Printf("%d\n", sum)
}
