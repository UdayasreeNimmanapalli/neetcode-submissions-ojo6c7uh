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
	inOrder(root, &arr)
	return arr[k-1]
}

func inOrder(node *TreeNode, arr *[]int){
	if node == nil{
		return
	}

	inOrder(node.Left, arr)

	*arr = append(*arr, node.Val)

	inOrder(node.Right, arr)
}
