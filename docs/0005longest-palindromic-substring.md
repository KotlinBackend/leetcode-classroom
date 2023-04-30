# 5. [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)
## 題目
Given a string s, return the longest palindromic substring in s.

Example 1:
```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```
Example 2:
```
Input: s = "cbbd"
Output: "bb"
``` 

## 題目大意：
- 給定一個字符串s，返回s中最長的回文子字符串。

## 解題思路：

一般而言，尋找回文字符串的問題都可以用動態規劃（dynamic programming）解決。首先，考慮最基本的情況，當s只有一個字符時，它是回文的。接下來，如果s的長度為2，那麼當且僅當兩個字符相等時，s才是回文的。對於s的長度大於2的情況，如果它是回文的，那麼將它的第一個字符和最後一個字符刪除後，剩餘的子字符串還是回文的。

因此，我們可以建立一個二維的bool類型的數組dp，其中dp[i][j]表示s的子串從索引i到索引j是否為回文字符串。由於只有子串是回文的時候才能進一步擴展，因此對於dp[i][j]為true的情況，只有當dp[i+1][j-1]也為true時，才能進一步擴展。最後，只需遍歷dp數組，找到最長的回文子字符串即可。

## Code
```go
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
        for i := 0; i + l < n; i++ {
            j := i + l
            if l == 0 {
                dp[i][j] = true
            } else if l == 1 {
                dp[i][j] = s[i] == s[j]
            } else {
                dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
            }

            // 如果從i到j的子串是回文的，並且比目前已知的最長回文子串更長，就更新答案
            if dp[i][j] && l + 1 > len(ans) {
                ans = s[i:i+l+1]
            }
        }
    }
    return ans
}
```

### [home](https://github.com/KotlinBackend/leetcode-classroom)