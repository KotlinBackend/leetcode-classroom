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
在陣列中找到兩個數之和等於給定值的數字，結果返回兩個數字在陣列中的索引。


## 解題思路
這題的最優解法時間複雜度為 O(n)。可以按照順序掃描陣列，對於每一個元素，在map中查找能夠與其組成給定值的另一半數字，如果找到了，直接返回這兩個數字在陣列中的索引即可。如果找不到，就把這個數字存入map中，等待掃描到“另一半”數字的時候，再從map中取出它並返回結果。換句話說，可以利用map的查找效率快速定位符合條件的數字對。

## Code
```go
package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

```