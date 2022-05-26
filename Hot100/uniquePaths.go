package main

// https://leetcode.cn/problems/unique-paths/

// DP equation: dp[i][j]=dp[i-1][j]+dp[i][j-1]
func uniquePathsDP(m int, n int) int {
	dp := make([][]int, m)
	for k := 0; k < m; k++ {
		dp[k] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// Use combination formule: The total numbers of action is m-1+n-1=m+n-2 which
// has m-1 left so the taotal combiantions obtain (n*...*m+n-2)div(m-1)!
func uniquePaths(m int, n int) int {
	var a, b int64
	a, b = 1, 1
	for i := n; i < m+n-1; i++ {
		a = a * int64(i)
	}
	for j := 1; j < m; j++ {
		b = b * int64(j)
	}
	return int(a / b)
	// return int(a/ b
}
