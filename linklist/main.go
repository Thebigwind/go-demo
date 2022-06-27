package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func main() {
	length := 5
	head := &node{0, nil}
	head = createLink(head, length)
	printLink(head)
	//reversedHead := reverseLink(head)
	reversedHead := reverseNode(head)
	printLink(reversedHead)

	fmt.Printf("--------\n")
	mid1 := lookMid(head)
	fmt.Println(mid1.data)

	mid2 := lookMid(reversedHead)
	fmt.Println(mid2.data)
}

func createLink(head *node, length int) *node {
	if length <= 0 {
		return head
	}
	for i := length - 1; i > 0; i-- {
		p := &node{i, nil}
		p.next = head.next
		head.next = p
	}
	return head
}

func printLink(head *node) {
	for p := head; p != nil; p = p.next {
		fmt.Print(p.data)
		fmt.Print(" ")
	}
	fmt.Println()
}

func reverseLink(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var reversedHead *node
	var p *node

	reversedHead = head
	head = head.next
	reversedHead.next = nil
	p = head.next
	for head != nil {
		head.next = reversedHead
		reversedHead = head
		head = p
		if p != nil {
			p = p.next
		}
	}
	return reversedHead
}

//求链表的中间节点
func lookMid(mid *node) *node {
	//求中间节点，快慢指针
	low := mid
	fast := mid
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		low = low.next
	}
	fmt.Println("中间节点的值为：%d", low.data)
	return low

}

//删除链表的倒数第n个节点

func reverseNode(head *node) *node {
	//  先声明两个变量
	//  前一个节点
	var preNode *node //= nil
	//  后一个节点
	var nextNode *node //  = new(node)
	fmt.Println(nextNode)
	//nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.next
		//  将头节点指向前一个节点
		head.next = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}

func rever(head *node) *node {
	var pre *node
	var next *node
	for head != nil {
		next = head.next //先保存head.next值
		head.next = pre  // 断开，让head.next指向pre,即nil
		pre = head       //更新pre值为head
		head = next      //更新head值为之前保存的heaad.next值
	}
	return pre
}

func rever2(head *node) *node {
	var pre *node = nil
	var next *node
	if head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock

*/
func maxProfit(prices []int) int {
	res := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > res {
			res = prices[i] - minPrice
		}
	}
	return res
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

type ListNode struct {
	Val  int
	Next *ListNode
}

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
