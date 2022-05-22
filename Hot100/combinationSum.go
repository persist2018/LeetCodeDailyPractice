package main

func combinationSum(candidates []int, target int) (ans [][]int) {
	sum := 0
	combinations := make([]int, 0)

	var backTrace func(Index int)
	backTrace = func(Index int) {
		if sum == target {
			ans = append(ans, combinations)
			return
		}
		if sum > target {
			return
		}
		for i := Index; i < len(candidates); i++ {
			sum = sum + candidates[i]
			combinations = append(combinations, candidates[i])
			backTrace(Index)
			combinations = combinations[:len(combinations)-1]
			sum = sum - candidates[i]
		}
	}
	backTrace(0)
	return ans
}

// func dfs(nums, combinations []int, target, index int) (c []int, ans [][]int) {
// 	fmt.Println("combinations:", combinations, "; target:", target, "; index:", index)
// 	if len(nums) == 0 || index == len(nums) || nums[index] > target {
// 		return
// 	}
// 	if nums[index] == target {
// 		combinations = append(combinations, nums[index])
// 		ans = append(ans, combinations)
// 		return
// 	}
// 	// skip this num
// 	combinations, ans = dfs(nums, combinations, target, index+1)
// 	// use this num
// 	// The question indicates the duration od nums[i] is [1,200]
// 	if target-nums[index] > 0 {
// 		combinations = append(combinations, nums[index])
// 		_, ans = dfs(nums, combinations, target-nums[index], index)
// 	}
// 	return
// }
