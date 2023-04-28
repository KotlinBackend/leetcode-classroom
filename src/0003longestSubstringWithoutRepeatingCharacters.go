package leetcode

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
