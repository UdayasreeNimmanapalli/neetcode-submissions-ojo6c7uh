/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    var maxVal = root.Val

	var dfs func(*TreeNode)int
	dfs = func(node *TreeNode)int{
		if node ==nil{
			return 0
		}
		leftMax := dfs(node.Left)
		rightMax := dfs(node.Right)
		leftMax = max(leftMax,0)
		rightMax = max(rightMax,0)
		if maxVal<(node.Val+leftMax+rightMax){
			maxVal = node.Val+leftMax+rightMax
		}
		return node.Val+max(leftMax, rightMax)
	}
	dfs(root)
	return maxVal
}
