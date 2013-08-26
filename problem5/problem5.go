package main

import "fmt"

func isDivisibleBy20Fact(num int) bool {
	divisorsArr := [...]int{20, 19, 18, 17, 16, 15, 14, 13, 11}
	for _, divisor := range divisorsArr {
		if num%divisor != 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 20; ; i += 20 {
		if isDivisibleBy20Fact(i) {
			fmt.Println(i)
			break
		}
	}
}
