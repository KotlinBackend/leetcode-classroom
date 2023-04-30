package leetcode

func longestPalindrome(s string) string {
	n := len(s)
	// 建立一個二維的bool類型數組dp，用來保存s的子串從索引i到索引j是否為回文字符串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	ans := ""
	// 枚舉字串的長度l
	for l := 0; l < n; l++ {
		// 枚舉子串的起始位置i，終止位置j
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}

			// 如果從i到j的子串是回文的，並且比目前已知的最長回文子串更長，就更新答案
			if dp[i][j] && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}
