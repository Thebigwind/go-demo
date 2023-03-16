package leetcode

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
