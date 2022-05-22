package main

// https://leetcode.cn/problems/permutations/

func permute(nums []int) (ans [][]int) {
	path := make([]int, 0)
	isVisited := make([]bool, len(nums))

	var backTrace func()
	backTrace = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}
		for i, num := range nums {
			if isVisited[i] {
				continue
			}
			path = append(path, num)
			isVisited[i] = true
			backTrace()
			isVisited[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrace()
	return
}
