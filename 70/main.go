package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(5))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n - 1) + climbStairs(n - 2)
}

func climbStairs1(n int) int {
	q, p, r := 0, 0, 1

	for i := 0; i < n; i++ {
		q = p
		p = r
		r = q + p
	}

	return r
}
