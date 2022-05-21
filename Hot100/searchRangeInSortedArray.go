package main

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	leftBound := searchLeftBoundInSortedArray(nums, 0, len(nums)-1, target)
	rightBound := searchRightBoundInSortedArray(nums, 0, len(nums)-1, target)
	if leftBound > len(nums) || rightBound > len(nums) {
		return []int{-1, -1}
	}
	if nums[leftBound] == target && nums[rightBound] == target {
		return []int{leftBound, rightBound}
	}
	return []int{-1, -1}
}

// change tradtioinal binary search
func searchLeftBoundInSortedArray(nums []int, left, right, target int) (index int) {
	var mid int
	for left <= right {
		mid = left - (left-right)/2
		if nums[mid] == target {
			right = mid - 1
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return left
}

func searchRightBoundInSortedArray(nums []int, left, right, target int) (index int) {
	var mid int
	for left <= right {
		mid = left - (left-right)/2
		if nums[mid] == target {
			left = mid + 1
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return right
}
