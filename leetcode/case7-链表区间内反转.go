package leetcode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//保存头节点
	pre := &ListNode{Next: head}
	newhead := pre
	// 1～m,顺序不变
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	//m~m+n,反转
	curr := pre.Next
	for i := m; i < n; i++ {
		t := curr.Next
		curr.Next = t.Next
		t.Next = pre.Next
		pre.Next = t
	}
	return newhead.Next
}
