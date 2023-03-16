package leetcode

type node struct {
	data int
	next *node
}

///////////// 给定一个排序链表，删除全部重复的元素，使得每一个元素只出现一次。////////////////
func deleteDuplicates(head *node) *node {
	curr := head
	for curr != nil && curr.next != nil {
		if curr.data == curr.next.data {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
	return head
}

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	curr := head
	record := map[int]int{}

	for curr != nil {
		record[curr.Val] = record[curr.Val] + 1 //记录每个元素的数量
		curr = curr.Next
	}

	curr = head
	for curr != nil && curr.Next != nil {
		if record[curr.Next.Val] > 1 {
			//需要删除
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	//判断head是否
	if head != nil && record[head.Val] > 1 {
		head = head.Next
	}

	return head
}

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
