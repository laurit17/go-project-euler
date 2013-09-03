package main

import (
	"fmt"
	"github.com/laurit17/go_utils/math_utils"
)

func GeneratePythagoreanTriple(m, n int) (a, b, c int) {
	a = math_utils.Pow(m, 2) - math_utils.Pow(n, 2)
	b = 2 * m * n
	c = math_utils.Pow(m, 2) + math_utils.Pow(n, 2)
	return
}

func main() {
	for m := 1; ; m++ {
		for n := 1; n < m; n++ {
			a, b, c := GeneratePythagoreanTriple(m, n)
			if a+b+c == 1000 {
				fmt.Println(a * b * c)
				return
			}
		}

	}
}
