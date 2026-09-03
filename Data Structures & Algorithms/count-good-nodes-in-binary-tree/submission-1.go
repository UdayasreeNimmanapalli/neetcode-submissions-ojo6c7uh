/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}
	var res int
	var dfs func(node *TreeNode, maxVal int)int
	dfs = func(node *TreeNode, maxVal int)int{
		if node == nil{
			return 0
		}
		if node.Val>=maxVal{
			res++
			maxVal = node.Val
		}
		dfs(node.Left, maxVal)
		dfs(node.Right, maxVal)

		return res
	}
	return dfs(root, root.Val)
}
