package main

// https://leetcode.cn/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	numMap := map[int]bool{}
	longest := 0
	for _, num := range nums {
		numMap[num] = true
	}

	for num := range numMap {
		if !numMap[num-1] {
			currentNum := num
			currentLength := 1

			for numMap[currentNum+1] {
				currentNum++
				currentLength++
			}

			if longest < currentLength {
				longest = currentLength
			}
		}
	}
	return longest
}
