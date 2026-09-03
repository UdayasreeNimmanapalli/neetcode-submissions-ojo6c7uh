/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int,0)
    if root == nil{
		return nil
	}
	var queue = make([]*TreeNode,0)
	queue = append(queue, root)
	for len(queue)>0{
		qlen:= len(queue)
		var temp = make([]int,0)
		for i:=0;i<qlen;i++{
			top := queue[0]
			queue = queue[1:]
			if top.Left!=nil{
				queue = append(queue, top.Left)
			}
			if top.Right != nil{
				queue = append(queue, top.Right)
			}
			temp = append(temp, top.Val)
		}
		res = append(res, temp)
	}
	return res
}
