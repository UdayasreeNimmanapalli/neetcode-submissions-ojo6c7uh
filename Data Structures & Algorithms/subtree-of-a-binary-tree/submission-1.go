/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
		return false
	}
	if subRoot == nil{
		return true
	}
	if isSametree(root, subRoot){
		return true
	}
	return isSubtree(root.Left,subRoot) || isSubtree(root.Right, subRoot)
}

func isSametree(root *TreeNode, subRoot *TreeNode)bool{
	if root==nil && subRoot == nil{
		return true
	}
	if (root!=nil && subRoot ==nil) || (root == nil && subRoot != nil){
		return false
	}
	if root!=nil && subRoot!=nil && root.Val==subRoot.Val{
		return isSametree(root.Left, subRoot.Left) && isSametree(root.Right, subRoot.Right)
	}
	return false
}
