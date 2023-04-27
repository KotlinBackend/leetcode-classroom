package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

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
