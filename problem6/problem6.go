package main

import (
	"fmt"
	"github.com/laurit17/go_utils/project_euler_utils"
)

func main() {
	sum := project_euler_utils.TriangleSum(100)
	result := sum*sum - project_euler_utils.SumOfSquares(100)
	fmt.Printf("%d\n", result)
}
