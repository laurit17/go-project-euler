package main

import (
  "fmt"
  "github.com/laurit17/go_utils/project_euler_utils"
)

func main() {
  triangSum := 0
  for i := 0;;i++ {
    triangSum = project_euler_utils.TriangleSum(i)
    if project_euler_utils.NumDivisors(triangSum) > 500 {
      fmt.Println(triangSum)
      return
    }

  }
}
