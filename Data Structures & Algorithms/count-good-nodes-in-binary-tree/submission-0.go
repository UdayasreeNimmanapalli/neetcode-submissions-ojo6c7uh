/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return goodNodesHelper(root, root.Val)
}

func goodNodesHelper(root *TreeNode, maxSofar int) int{
	if root == nil{
		return 0
	}
	var count = 0
	if root.Val >=maxSofar{
		count++
	}
	if root.Val > maxSofar{
		maxSofar = root.Val
	}
	return count+goodNodesHelper(root.Left, maxSofar)+goodNodesHelper(root.Right, maxSofar)
}
