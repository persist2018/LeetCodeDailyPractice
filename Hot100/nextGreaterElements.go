package main

// https://leetcode.cn/problems/next-greater-element-ii/
func nextGreaterElements(nums []int) []int {
	length := len(nums)
	if length == 1 {
		return []int{-1}
	}
	ans := make([]int, 0)
	doubleNums := make([]int, 2*length)
	for i := 0; i < 2*length; i++ {
		if i <= length-1 {
			doubleNums[i] = nums[i]
		} else {
			doubleNums[i] = nums[i-length]
		}
	}
	// fmt.Println(doubleNums)
	for j := 0; j < length; j++ {
		for k := j; k < j+length+1; k++ {
			if doubleNums[k] > doubleNums[j] {
				ans = append(ans, doubleNums[k])
				break
			}
			if k == j+length {
				ans = append(ans, -1)
			}
		}
	}
	return ans

}
