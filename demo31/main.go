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
