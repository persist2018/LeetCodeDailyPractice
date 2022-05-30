package main

import (
	"fmt"
)

// https://leetcode.cn/problems/palindrome-number/
func isPalindrome(x int) bool {
	xStr := fmt.Sprint(x)
	length := len(xStr)
	left, right := 0, length-1
	for left < right {
		if xStr[left] != xStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
