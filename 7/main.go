package main

import "fmt"

func reverse(n int) int {
	res := 0

	for n != 0 {
		res = res * 10 + n % 10
		n = n / 10
	}

	if res > 1 << 31 - 1 || res < -(1 << 31) {
		return 0
	}

	return res
}

func main() {
	fmt.Println(reverse(-3210))
}
