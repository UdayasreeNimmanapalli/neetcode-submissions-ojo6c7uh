/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter = 0
   if root == nil{
		return 0
   } 
   lengthOfBinaryTree(root, &diameter)
   return diameter
}

func lengthOfBinaryTree(root *TreeNode, diameter *int)int{
	if root == nil{
		return 0
	}
	left := lengthOfBinaryTree(root.Left, diameter)
	right := lengthOfBinaryTree(root.Right, diameter)
	d:=left+right
	if d>*diameter{
		*diameter = d
	}
	return 1+max(left, right)
}