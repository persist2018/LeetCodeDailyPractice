package main

// https://leetcode.cn/problems/minimum-window-substring/

// This problem supposed there only exist at most one possible solution.
func minWindow(s string, t string) string {
	lengthT := len(t)
	lengthS := len(s)
	if s == "" || t == "" || lengthS < lengthT {
		return ""
	}
	left, right, distance := 0, 0, 0
	tMap, windowMap := make(map[rune]int, 0), make(map[rune]int, 0)
	ansLeft, minLen := 0, lengthS+1
	for _, tChar := range t {
		tMap[tChar]++
	}

	// find a solution
	for right < lengthS {
		if tMap[rune(s[right])] == 0 {
			right++
			continue
		}
		if tMap[rune(s[right])] > windowMap[rune(s[right])] {
			distance++
		}
		windowMap[rune(s[right])]++
		right++
		// fmt.Println(windowMap)
		for distance == lengthT {
			// fmt.Println(distance)
			if right-left < minLen {
				minLen = right - left
				ansLeft = left
			}
			if tMap[rune(s[left])] == 0 {
				left++
				continue
			}
			if tMap[rune(s[left])] == windowMap[rune(s[left])] {
				distance--
			}
			windowMap[rune(s[left])]--
			left++
		}
	}
	if minLen == lengthS+1 {
		return ""
	}
	return s[ansLeft : ansLeft+minLen]
}

func checkCoverage(input, window map[rune]int) bool {
	for k, v := range input {
		if window[k] < v {
			return false
		}
	}
	return true
}
