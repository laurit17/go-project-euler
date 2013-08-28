package main

import (
	"fmt"
	"os"
	"strings"
)

func MaxProductInString(s string) int {
	runes := []rune(s)
	max := 0
	for i := 0; i < len(runes)-4; i++ {
		current_slice := runes[i : i+5]
		product := 1
		for _, value := range current_slice {
			product *= RuneToInt32(value)
		}
		if product > max {
			max = product
		}
	}
	return max
}

func RuneToInt32(r rune) int {
	return int(r - '0')
}

func main() {
	file, err := os.Open("number.txt")
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
	str = strings.Replace(str, "\n", "", -1)

	fmt.Println(MaxProductInString(str))
}
