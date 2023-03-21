package leetcode

// https://leetcode.cn/problems/add-two-numbers/
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	//定义一个尾结点，或者可以理解为临时节点
	var tail *ListNode
	//余数carry
	carry := 0
	//依次遍历两个链表，只要元素不为空就进行下一步
	for l1 != nil || l2 != nil {
		//定义两个变量存储各个节点的值
		n1, n2 := 0, 0
		//从第一个链表开始
		if l1 != nil {
			//把每个节点的值赋给n1
			n1 = l1.Val
			//节点后移
			l1 = l1.Next
		}
		//l2同上
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//此时是两个链表第一个元素的和 + 余数
		sum := n1 + n2 + carry
		//sum%10是节点的当前值，如果是10,取余后当前节点值为0，sum/10是求十位的那个数
		sum, carry = sum%10, sum/10
		//此时申请一个新的链表存储两个链表的和
		if head == nil {
			//申请新的链表
			head = &ListNode{Val: sum}
			//这一步是为了保持头结点不变的情况下指针可以右移，所以说tail相当于临时节点，理解成尾节点也可以，因
			//为此时新链表中只有一个节点，所以头结点和尾结点都指向同一个元素。
			tail = head
		} else {
			//第二个节点后开始逐渐往尾结点增加元素
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	//把最后一位的余数加到链表最后。
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = new(ListNode) // 新建一个节点存储计算结果
		curr = curr.Next          // 将计算结果连接成链表
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry /= 10
	}
	return dummy.Next
}

//链接：https://leetcode.cn/problems/add-two-numbers/solution/go-chao-jian-ji-by-xilepeng/
