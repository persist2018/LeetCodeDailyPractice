package main

func searchRotatedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] != target {
			return -1
		}
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left - (left-right)/2
		// fmt.Println(nums[mid])
		if nums[mid] == target {
			return mid
		} else {
			if nums[left] <= nums[mid] {
				if target < nums[mid] && target >= nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				if target > nums[mid] && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return -1
}
