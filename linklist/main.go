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
