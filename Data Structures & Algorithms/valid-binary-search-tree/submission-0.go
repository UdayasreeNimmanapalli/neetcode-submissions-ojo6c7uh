/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64,math.MaxInt64)
}

func isValidBSTHelper(node *TreeNode, min, max int)bool{
	if node == nil{
		return true
	}
	if node.Val<= min || node.Val>=max{
		return false
	}
	return isValidBSTHelper(node.Left, min, node.Val) && isValidBSTHelper(node.Right, node.Val, max)
}