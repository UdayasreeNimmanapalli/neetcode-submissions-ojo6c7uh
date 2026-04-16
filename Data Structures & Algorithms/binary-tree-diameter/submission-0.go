/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter = 0
    calculateHeight(root, &maxDiameter)
	return maxDiameter
}

func calculateHeight(node *TreeNode, diameter *int)int{
	if node ==nil{
		return 0
	}
	maxLeft := calculateHeight(node.Left, diameter)
	maxRight := calculateHeight(node.Right,diameter)
	d := maxLeft+maxRight
	if d>*diameter{
		*diameter = d
	}
	return 1+max(maxLeft, maxRight)
}
