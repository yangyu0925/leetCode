package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	res := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			res = res * 10 + int(v - '0')
		} else if i == 0 && str[i] == '-' {
			sign = -1
		} else if i == 0 && str[i] == '+' {
			sign = 1
		} else {
			break
		}

		if res > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}

	return sign * res
}

func main() {
	fmt.Println(myAtoi("798hello world"));
}

