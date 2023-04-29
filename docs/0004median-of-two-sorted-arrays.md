# 4. [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)
## 題目
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).


Example 1:
```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```
Example 2:
```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

## 題目大意
- 給定兩個已排序的數組 nums1 和 nums2，大小分別為 m 和 n，求其合併後的中位數。


## 解題思路
- 首先我們需要先將兩個數組合併成一個有序的數組，可以使用雙指針來實現。然後再根據數組的長度奇偶性，求出中位數。
- 若數組長度為奇數，則中位數為排序後的數組中間位置的數；若數組長度為偶數，則中位數為排序後的數組中間兩個數的平均值。
- 由於題目要求算法的時間複雜度為 O(log(m+n))，所以我們需要使用二分查找來實現。具體來說，對於長度較小的數組，我們可以使用二分查找來找到其分界點，然後在另一個數組中二分查找對應的位置。

## Code
```go
package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m // 確保 nums2 長度比 nums1 長
	}
	iMin, iMax, halfLen := 0, m, (m+n+1)/2 // 設置二分查找範圍和中間值
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < m && nums2[j-1] > nums1[i] { // i 偏小，需要增大
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] { // i 偏大，需要減小
			iMax = i - 1
		} else { // i 小於等於 m 且 j 小於等於 n，找到合適的 i
			var maxLeft float64
			if i == 0 { // nums1 中所有元素都大於 nums2 的最大值，直接從 nums2 中取最大值
				maxLeft = float64(nums2[j-1])
			} else if j == 0 { // nums2 中所有元素都大於 nums1 的最大值，直接從 nums1 中取最大值
				maxLeft = float64(nums1[i-1])
			} else { // 取 nums1[i-1] 和 nums2[j-1] 中的較大值
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 { // 數組總長度為奇數，中位數即為 maxLeft
				return maxLeft
			}
			var minRight float64
			if i == m { // nums1 中所有元素都小於 nums2 的最小值，直接從 nums2 中取最小值
				minRight = float64(nums2[j])
			} else if j == n { // nums2 中所有元素都小於 nums1 的最小值，直接從 nums1 中取最小值
				minRight = float64(nums1[i])
			} else { // 取 nums1[i] 和 nums2[j] 中的較小值
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (maxLeft + minRight) / 2 // 數組總長度為偶數，中位數為左側最大值與右側最小值的平均值
		}
	}
	return 0.0 // 當 nums1 和 nums2 為空數組時，返回 0
}

```