package main

func longestPalindrome(s string) string {
	length := len(s)
	maxLen := 1
	begin := 0
	// init dp status
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	// when lengthOfPalindromeh=1
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}

	for lengthOfPalindrome := 2; lengthOfPalindrome <= length; lengthOfPalindrome++ {
		for i := 0; i < length; i++ {
			j := lengthOfPalindrome + i - 1
			if j >= length {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i > 3 {
					dp[i][j] = dp[i+1][j-1]
				} else {
					dp[i][j] = true
				}
			}
		}
	}

}
