package main

import "fmt"

func main() {
	arr := []int{4, 6, 2, 7, 34, 8, 9, 12, 4, 23, 75, 88, 2, 0, 24}
	fmt.Println(qsort(arr))
}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	//已最右侧的元素为分区点
	for i := range a {
		if a[i] < a[right] {
			//将a[i]放到左侧分区
			a[left], a[i] = a[i], a[left]
			left++ //左侧分区+1
		}
	}
	//将分区点放到合适的位置
	a[left], a[right] = a[right], a[left]
	//递归
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

type node struct {
	data int
	next *node
}

func reverse(head *node) *node {
	var pre *node = nil
	var next *node
	if head != nil {
		next = head.next //保存头结点的下一个结点
		head.next = pre  //断开head.next,改为指向pre
		pre = head       //更新 pre 值为head
		head = next      // 更新 head 值为原来head.next保存到next里的值
	}
	return pre
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

func hasCycle(head *node) bool {
	seen := map[*node]struct{}{} // 开一个map记录该节点是否已经遍历过，值记录节点索引
	for head != nil {
		if _, ok := seen[head]; ok { // 该节点遍历过，形成了环
			return true
		}
		seen[head] = struct{}{} /// 记录该节点已经遍历过
		head = head.next
	}
	return false
}
