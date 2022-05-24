package main

// https://leetcode.cn/problems/minimum-path-sum/

// DP: dp[i][j] =min(dp[i-1][j],dp[i][j-1])+grid[i][j]
func minPathSumDP(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for k := 0; k < m; k++ {
		dp[k] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}
