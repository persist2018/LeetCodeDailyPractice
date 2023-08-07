package main

// https://leetcode.cn/problems/jump-game/
// This problem can be solved by greedy algorithm
func canJump(nums []int) bool {
	maxMost := 0

	for i := 0; i < len(nums); i++ {
		if maxMost >= i {
			maxMost = max(maxMost, i+nums[i])
			if i == len(nums)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
