/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	dfs(root,&arr)
	return arr[k-1]
}

func dfs(node *TreeNode, arr *[]int){
	if node == nil{
		return
	}
	dfs(node.Left, arr)
	*arr = append(*arr, node.Val)
	dfs(node.Right, arr)
}
