package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freg [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right + 1 < len(s) && freg[s[right + 1] - 'a'] == 0 {
			freg[s[right + 1] - 'a']++
			right++
		} else {
			freg[s[left] - 'a']++
			left++
		}
		result = max(result, right - left + 1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
