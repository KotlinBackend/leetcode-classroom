# 6. [Zigzag Conversion](https://leetcode.com/problems/zigzag-conversion/)
## 題目
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
 

Example 1:
```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```
Example 2:
```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
```
Example 3:
```
Input: s = "A", numRows = 1
Output: "A"
```

## 題目大意：
- 給定一個字符串 s 和一個正整數 numRows，將 s 轉換成一個 "之" 字形，每行由左斜向右斜，再由右斜向左斜排列。返回按行讀取的新字符串。

## 解題思路：
建立一個長度為 numRows 的列表，每個元素代表之字形的一行，並遍歷原字符串 s，按照之字形的順序將每個字符添加到對應的行。遍歷完畢後，將所有行拼接起來即為結果。

## Code
```go
func convert(s string, numRows int) string {
    // 特判：只有一行的情況，直接返回原字符串
    if numRows == 1 {
        return s
    }

    // 創建 numRows 個空字符串，用於存儲各行字符
    rows := make([]string, numRows)

    // curRow：當前處理的行數；goingDown：是否往下走（即是否需要向下遍歷）
    curRow, goingDown := 0, false

    // 遍歷字符串中的每個字符
    for _, ch := range s {
        // 將當前字符添加到對應的行中
        rows[curRow] += string(ch)
        // 如果當前行位於第一行或最後一行，則需要改變方向
        if curRow == 0 || curRow == numRows-1 {
            goingDown = !goingDown
        }
        // 根據方向改變當前行
        if goingDown {
            curRow++
        } else {
            curRow--
        }
    }

    // 拼接所有行，得到結果
    res := ""
    for _, row := range rows {
        res += row
    }
    return res
}

```