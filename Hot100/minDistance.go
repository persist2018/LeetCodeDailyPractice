package main

// https://leetcode.cn/problems/edit-distance/
// DP: dp[i][j] =1+ minThree(dp[i-1][j], dp[i][j-1],dp[i-1][j-1]-1)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// handle ome of thre input is ""
	if n*m == 0 {
		return n + m
	}
	// length is n+1 and m+1
	dp := make([][]int, m+1)
	for k := 0; k <= m; k++ {
		dp[k] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else if i == 0 && j != 0 {
				dp[i][j] = j
			} else if j == 0 && i != 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = 1 + minThree(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]-1)
				} else {
					dp[i][j] = 1 + minThree(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				}
			}
		}
	}
	return dp[m][n]

}
