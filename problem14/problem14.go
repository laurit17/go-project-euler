package main

import "fmt"

func collatz(n int64) int {

	sum := 1
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		sum += 1
	}
	return sum
}

func collatzRecursive(n int64) int {

	switch {
	case n == 1:
		return 1
	case n%2 == 0:
		return 1 + collatzRecursive(n/2)
	case n%2 == 1:
		return 1 + collatzRecursive(3*n+1)

	}
  return -1 
}

func collatzWithCaching(length int) func(int64) int {
  cache := make([]int, length + 1)
  var temp func(num int64) int
  i64len := int64(length)
  temp = func(num int64) int {
    if num == 1 {
      return 1
    }
    if num <= i64len && cache[num] != 0 {
      return cache[num]
    }
    var result int
    if num % 2 == 0 {
      result = temp(num/2) + 1
    } else {
      result = temp(3*num + 1) + 1
    }

    if num <= i64len {
      cache[num] = result
    }
    return result
  }

  return temp

}

func main() {
  collatzFast := collatzWithCaching(100000)
	max, curr, longestChain := 0, 0, 0
	for i := 1; i < 1000000; i++ {
		curr = collatzFast(int64(i))
		if curr > max {
      longestChain = i
			max = curr
		}
	}

	fmt.Printf("%v has the longest Collatz chain\n", longestChain)
}
