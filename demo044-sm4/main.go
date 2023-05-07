package main

import "fmt"

//sm4
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	fmt.Println(sumNumbers(root))
}

//122 + 125 + + 13
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
func sumNumbers(root *TreeNode) int {
	var res []int
	var num, sum int
	dfs(root, num, &res)
	for i := 0; i < len(res); i++ {
		sum += res[i]
	}
	return sum
}

func dfs(root *TreeNode, num int, res *[]int) {
	if root == nil {
		return
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res = append(*res, num)
	}
	dfs(root.Left, num, res)
	dfs(root.Right, num, res)
}
