package main

// https://leetcode.cn/problems/trapping-rain-water/

func trappingRainDP(height []int) (ans int) {
	length := len(height)
	if length == 0 {
		return 0
	}
	leftMax := make([]int, length)
	rightMax := make([]int, length)
	// leftMax[i]=max(leftMax[i-1],hight[i])
	for i := 0; i < length; i++ {
		if i == 0 {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = max(height[i], leftMax[i-1])
		}
	}
	// rightMax[i]=max(rightMax[i+1],hight[i])
	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = max(height[i], rightMax[i+1])
		}
	}
	// rain[i]=min(leftMax[i],rightMax[i])-height[i]

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
}

func trappingRainTwoPointers(height []int) (ans int) {
	length := len(height)
	if length == 0 {
		return
	}
	left, right, leftMax, rightMax := 0, length-1, height[0], height[length-1]
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right++
		}
	}
	return
}
