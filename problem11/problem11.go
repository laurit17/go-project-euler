package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func searchRowsForBiggestProduct(grid [][]int) int {
	max, product := 0, 1
	for _, row := range grid {
		for i := 0; i < len(row)-3; i++ {
			product = row[i] * row[i+1] * row[i+2] * row[1+3]
			if product > max {
				max = product
			}
		}
	}

	return max

}

func searchColumnsForBiggestProduct(grid [][]int) int {
	max, product := 0, 1
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid)-3; j++ {
			product = grid[i][j] * grid[i][j+1] * grid[i][j+2] * grid[i][j+3]
			if product > max {
				max = product
			}
		}
	}

	return max
}

func searchMajorDiagsForBiggestProduct(grid [][]int) int {
	max, product := 0, 1

	for i := 0; i < len(grid)-3; i++ {
		for j := 0; j < len(grid[0])-3; j++ {
			product = grid[i][j] * grid[i+1][j+1] * grid[i+2][j+2] * grid[i+3][j+3]
			if product > max {
				max = product
			}
		}
	}

	return max
}

func searchMinorDiagsForBiggestProduct(grid [][]int) int {
	max, product := 0, 1

	for i := 0; i < len(grid[0])-3; i++ {
		for j := 3; j < len(grid[0]); j++ {
			product = grid[i][j] * grid[i+1][j-1] * grid[i+2][j-2] * grid[i+3][j-3]
			if product > max {
				max = product
			}
		}
	}
	return max
}

func Max(nums []int) int {
	max := 0
	for _, value := range nums {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	file, err := os.Open("grid.txt")
	if err != nil {
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	lines := strings.Split(str, "\n")
	lines = lines[:len(lines)-1]
	grid := make([][]int, len(lines))
	for index, value := range lines {
		nums := strings.Split(strings.TrimSpace(value), " ")
		grid[index] = make([]int, len(nums))
		for inner_index, inner_value := range nums {
			grid[index][inner_index], err = strconv.Atoi(inner_value)
			if err != nil {
				return
			}
		}
	}

	products := []int{
		searchRowsForBiggestProduct(grid),
		searchColumnsForBiggestProduct(grid),
		searchMajorDiagsForBiggestProduct(grid),
		searchMinorDiagsForBiggestProduct(grid),
	}
	fmt.Println(Max(products))
}
