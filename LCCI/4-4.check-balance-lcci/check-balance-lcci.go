/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func depth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(depth(root.Left), depth(root.Right)) + 1

}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(depth(root.Left)-depth(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}