package main

import (
	"fmt"
	"github.com/laurit17/go_utils/math_utils"
	"os"
	"strconv"
	"strings"
)

func DivideIntoListsOf10(s string) []int64 {
	//we know that they are length 50 here
	smallerInts := make([]int64, len(s)/10)
	for i := 0; i < len(s)/10; i++ {
		value, err := strconv.Atoi(s[10*i : 10*(i+1)])
		if err != nil {
			return nil
		}
		smallerInts[i] = int64(value)
	}

	return smallerInts
}

func SumOverArray(arr []int64) int64 {
	sum := int64(0)
	for _, value := range arr {
		sum += value
	}
	return sum
}

func ChopTrailing10Digits(n int64) int64 {
	return n / int64(math_utils.Pow(10, 10))
}

func main() {
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println(err)
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
	brokenNumbers := make([][]int64, len(lines))
	for index, value := range lines {
		brokenNumbers[index] = DivideIntoListsOf10(value)
	}

	groupedByPlace := make([][]int64, 5)
	for index := range groupedByPlace {
		groupedByPlace[index] = make([]int64, len(brokenNumbers))
	}

	for index := range brokenNumbers {
		for i := 0; i < 5; i++ {
			groupedByPlace[i][index] = brokenNumbers[index][4-i]
		}
	}

	leadingDigits, sum := int64(0), int64(0)
	for _, value := range groupedByPlace {
		sum = SumOverArray(value) + leadingDigits
		leadingDigits = ChopTrailing10Digits(sum)
	}

	fmt.Println(strconv.FormatInt(sum, 10)[:10])

}
