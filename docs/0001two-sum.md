# 1. [Two Sum](https://leetcode.com/problems/two-sum/)
## 題目
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

### Example 1:
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```
### Example 2:
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```
### Example 3:
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

## 題目大意
- 給定一個整數陣列 nums 和一個整數目標值 target，請在陣列中找到兩個元素，使得它們的和等於目標值 target，並返回這兩個元素的索引。
- 可以假設每個輸入都只有一個解，且同一個元素不能使用兩次。可以以任何順序返回答案。


## 解題思路
- 這個問題可以使用 map 來解決，遍歷整個陣列，對於每個元素，判斷目標值與當前元素值之差是否在之前遍歷過的元素中出現過。如果差值存在於 map 中，則返回這兩個元素的索引，否則將當前元素的值和索引加入 map 中，繼續遍歷。
- 使用 map 可以在 O(1) 的時間複雜度內查找元素的索引，因此整個算法的時間複雜度為 O(n)。
- 具體的解題思路：
	1. 創建一個空的 map，用來存儲元素值和它們的索引。
	2. 遍歷整數陣列，對於每個元素，計算目標值與當前元素值之差，並在 map 中查找差值是否	存在。
	3. 如果差值存在於 map 中，則返回這兩個元素的索引。
	4. 否則，將當前元素值和它的索引加入 map 中，繼續遍歷。
	5. 如果找不到符合要求的元素，則返回空。

## Code
```go
package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // 創建一個空的 map，用來存儲元素值和它們的索引
	for i, num := range nums {
		another := target - num          // 計算目標值與當前元素值之差
		if index, ok := m[another]; ok { // 如果差值存在於 map 中，則返回這兩個元素的索引
			return []int{index, i}
		}
		m[num] = i // 否則，將當前元素值和它的索引加入 map 中
	}
	return nil // 如果找不到符合要求的元素，則返回空
}
```

### [home](https://github.com/KotlinBackend/leetcode-classroom)