/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	var res = make([]int,0)
	rightSideViewHelper(root, 0, &res)
	return res
}
func rightSideViewHelper(root *TreeNode,level int, res *[]int){
	if root == nil{
		return
	}
	if level == len(*res){
		*res = append(*res, root.Val)
	}
	rightSideViewHelper(root.Right, level+1, res)
	rightSideViewHelper(root.Left, level+1, res)
}
