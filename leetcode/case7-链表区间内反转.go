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

func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func reverseBetween2(head *ListNode, left, right int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	reverseLinkedList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

//链接：https://leetcode.cn/problems/reverse-linked-list-ii/solution/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/

func reverseBetween3(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}

	// 记住头节点
	res := &ListNode{}
	// leftPrev是left的前一个节点 用来和指定的反转链表合并
	leftPrev := res

	// 指定的反转链表
	var recv *ListNode
	// 指定的反转链表最后一个节点 也是left节点
	// 用来挂载head最后的剩余节点 凑成完整的返回结果
	var recvTailNode *ListNode

	// 记录循环次数
	num := 0

	for head != nil {
		num += 1

		// 反转链表
		next := head.Next
		head.Next = recv
		recv = head

		if num < left {
			leftPrev.Next = head
			leftPrev = leftPrev.Next
		}

		if num == left {
			recvTailNode = recv
			// 将left之前的节点从recv删除
			recv.Next = nil
		}

		if num == right {
			// 挂载剩余的链表
			recvTailNode.Next = next
			break
		}

		// 反转链表
		head = next
	}

	leftPrev.Next = recv
	return res.Next
}
