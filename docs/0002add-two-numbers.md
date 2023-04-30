# 2. [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)
## 題目
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

 
Example 1:
```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```
Example 2:
```
Input: l1 = [0], l2 = [0]
Output: [0]
```
Example 3:
```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

## 題目大意
- 給定兩個非空的連結串列，分別代表兩個非負整數。每個數位被反向存儲在連結串列中，每個節點包含一個數位。將兩個數相加，並將其作為連結串列返回。
- 假設兩個數字都不包含任何前導零，除了數字0本身。

## 解題思路
此問題需要對兩個連結串列進行加法運算，運算的過程和手算加法類似，需要從個位開始進行相加並考慮進位的情況。為了方便，可以新建一個節點存儲相加的結果，每次運算完當前節點後，將該節點插入到新建的節點之後。需要注意的是，如果兩個連結串列的長度不一致，那麼在運算時需要在短的連結串列中補零。

## Code
```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 建立虛擬節點 dummy，作為結果鏈表的頭節點
    dummy := &ListNode{}
    // curr 指向結果鏈表的最後一個節點
    curr := dummy
    // carry 存放進位值，初始值為 0
    carry := 0
    // 當 l1 或 l2 不為空時，進行遍歷
    for l1 != nil || l2 != nil {
        // val1 存放 l1 節點的值，若 l1 為空則預設為 0
        var val1, val2 int
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }
        // val2 存放 l2 節點的值，若 l2 為空則預設為 0
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }
        // sum 存放 val1、val2、carry 三者的和
        sum := val1 + val2 + carry
        // 計算新的進位值
        carry = sum / 10
        // 將當前位的結果存入結果鏈表
        curr.Next = &ListNode{Val: sum % 10}
        // 將 curr 指向結果鏈表的最後一個節點
        curr = curr.Next
    }
    // 如果最後 carry 不為 0，還要再在結果鏈表的最後加上一個節點
    if carry > 0 {
        curr.Next = &ListNode{Val: carry}
    }
    // 返回結果鏈表的頭節點的下一個節點（因為第一個節點是虛擬節點）
    return dummy.Next
}
```

### [home](https://github.com/KotlinBackend/leetcode-classroom)