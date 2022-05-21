package main

//
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	left := (length1 + length2 + 1) / 2
	right := (length1 + length2 + 2) / 2

	return float64(devide(nums1, nums2, left, 0, 0, length1-1, length2-1)+devide(nums1, nums2, right, 0, 0, length1-1, length2-1)) * 0.5
}
func devide(nums1, nums2 []int, k, start1, start2, end1, end2 int) int {
	len1 := -start1 + end1 + 1
	len2 := -start2 + end2 + 1
	if len1 > len2 {
		return devide(nums2, nums1, k, start2, start1, end2, end1)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}
	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1
	if nums1[i] > nums2[j] {
		return devide(nums1, nums2, k-(j-start2+1), start1, j+1, end1, end2)
	} else {
		return devide(nums1, nums2, k-(i-start1+1), i+1, start2, end1, end2)

	}

}
