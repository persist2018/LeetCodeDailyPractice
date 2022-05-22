package main

// https://leetcode.cn/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sum := 0
	combinations := make([]int, 0)

	var backTrace func(Index int)
	backTrace = func(Index int) {
		if sum == target {
			ans = append(ans, append([]int{}, combinations...))
			return
		}
		if sum > target {
			return
		}
		for i := Index; i < len(candidates); i++ {
			sum = sum + candidates[i]
			combinations = append(combinations, candidates[i])
			backTrace(i)
			combinations = combinations[:len(combinations)-1]
			sum = sum - candidates[i]
		}
	}
	backTrace(0)
	return ans
}
