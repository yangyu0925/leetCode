package main

import (
	"fmt"
	"strings"
)

func containsDuplicate(s string) string {
	ss := strings.Split(s, " ")
	for i, s := range ss {
		ss[i] = revers(s)
	}
	return strings.Join(ss, " ")
}

func revers(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes) - 1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func main() {
	fmt.Println(containsDuplicate("Let's take LeetCode contest"))
}