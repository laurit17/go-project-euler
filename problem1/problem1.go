package main

import (
	"fmt"
	"github.com/laurit17/go_utils/project_euler_utils"
)

func main() {
	threes_sum := project_euler_utils.DivisibleSumLessThan(1000, 3)
	fives_sum := project_euler_utils.DivisibleSumLessThan(1000, 5)
	fifteens_sum := project_euler_utils.DivisibleSumLessThan(1000, 15)
	sum := threes_sum + fives_sum - fifteens_sum
	fmt.Printf("%d\n", sum)
}
