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