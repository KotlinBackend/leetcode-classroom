# 3. [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
## 題目
Given a string s, find the length of the longest 
substring
without repeating characters.


Example 1:
```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```
Example 2:
```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```
Example 3:
```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```
 

## 題目大意
- 給定一個字符串 s，請找到該字符串中最長的不重複子串的長度。

## 解題思路
- 運用雙指針法解題，維護一個滑動窗口，窗口內沒有重複元素。當遇到重複元素時，右指針不斷右移直到窗口內沒有該重複元素，同時更新窗口大小，不斷更新最大子串的長度。

- 具體實現步驟：
  1. 宣告一個 map 變數來存儲元素，其中 key 是字符串中的字符，value 是該字符最近一次出現的位置。
  2. 宣告兩個指針，分別指向字符串 s 的左右邊界。
  3. 遍歷字符串 s，當右指針指向的元素已經存在於 map 中時，將左指針指向該元素最近出現的位置的下一個位置，同時將 map 中左指針位置和左指針位置之前的元素全部刪除。
  4. 在遍歷過程中維護最大不重複子串的長度，並返回最大子串的長度。

## Code
```go
func lengthOfLongestSubstring(s string) int {
	maxLen := 0                      // 最長無重複字元子串長度
	currLen := 0                     // 目前無重複字元子串長度
	start := 0                       // 目前無重複字元子串開始位置
	charIdxMap := make(map[byte]int) // 存儲目前已出現過的字元及其最新出現位置

	for i := 0; i < len(s); i++ {
		// 如果目前字元已出現在前面的子串中，並且出現位置在目前子串中，更新子串開始位置為上一次出現位置+1
		// 更新目前子串長度為從目前開始位置到目前字元位置的長度
		if lastIdx, ok := charIdxMap[s[i]]; ok && lastIdx >= start {
			start = lastIdx + 1
			currLen = i - start + 1
		} else {
			// 如果目前字元沒有出現在前面的子串中，目前子串長度加一
			currLen++
		}

		// 如果目前子串長度大於最長子串長度，更新最長子串長度
		if currLen > maxLen {
			maxLen = currLen
		}

		// 存儲目前字元及其最新出現位置
		charIdxMap[s[i]] = i
	}
	return maxLen
}

```

### [home](https://github.com/KotlinBackend/leetcode-classroom)