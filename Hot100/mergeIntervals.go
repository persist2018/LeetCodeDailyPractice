package main

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func mergeTwoIntervals(interval1, interval2 []int) [][]int {
	// intervalA[0]<intervalB[0]
	intervalA, intervalB := make([]int, 0), make([]int, 0)
	if interval1[0] < interval2[0] {
		copy(intervalA, interval1)
		copy(intervalB, interval2)
	} else {
		copy(intervalA, interval2)
		copy(intervalB, interval1)
	}
	if intervalA[1] < intervalB[0] {
		return [][]int{intervalA, intervalB}
	}
	if intervalB[0] < intervalA[1] && intervalB[1] > intervalA[1] {
		return [][]int{[]int{intervalA[0], intervalB[1]}}
	}
	if intervalB[1] < intervalA[1] {
		return [][]int{intervalA}
	}
	return [][]int{intervalA}
}
