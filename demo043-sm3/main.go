package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	fmt.Println(sumNumbers(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return int整型
 */

var path []int

func sumNumbers(root *TreeNode) int {
	path = make([]int, 0)
	dfs("", root)
	res := 0
	for _, v := range path {
		res += v
	}
	return res
}

func dfs(str string, root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		str += strconv.Itoa(root.Val)
		num, _ := strconv.Atoi(str)
		path = append(path, num)
		return
	}
	str += strconv.Itoa(root.Val)
	dfs(str, root.Left)
	dfs(str, root.Right)
}
