package leetcode

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
	// write code here
	if root == nil {
		return 0
	}
	ans := 0
	var helper func(parentVal int, node *TreeNode)
	helper = func(parentVal int, node *TreeNode) {
		parentVal *= 10
		nodeVal := node.Val
		parentVal += nodeVal
		if node.Left == nil && node.Right == nil {
			ans += parentVal
			return
		}
		if node.Left != nil {
			helper(parentVal, node.Left)
		}
		if node.Right != nil {
			helper(parentVal, node.Right)
		}
	}
	helper(0, root)
	return ans
}
