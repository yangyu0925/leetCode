package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
		fmt.Println(revertedNumber, x)

	}

	return x == revertedNumber || x == revertedNumber / 10
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)

	for i := range str {
		if str[len(str) - 1 - i] != str[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome1(12221))
}
