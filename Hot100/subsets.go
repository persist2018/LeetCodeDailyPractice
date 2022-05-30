package main

// https://leetcode.cn/problems/subsets/

func subsets(nums []int) (ans [][]int) {
	path := make([]int, 0)
	// isVisited := make([]bool, len(nums))
	var backTrace func(int)
	backTrace = func(start int) {
		if start == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			// this three lines also can be writen in:
			// ans = append(ans, append([]int(nil), path...))
			return
		}
		path = append(path, nums[start])
		backTrace(start + 1)
		path = path[:len(path)-1]
		backTrace(start + 1)
	}
	backTrace(0)
	return
}
