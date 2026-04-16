/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
    var q = make([]TreeNode,0)
	q = append(q,*root)
	count:=0
	for len(q)>0{
		qlen := len(q)
		for i:=0;i<qlen;i++{
			top := q[0]
			q=q[1:]
			if top.Left!=nil{
				q = append(q, *top.Left)
			}
			if top.Right!=nil{
				q = append(q, *top.Right)
			}
		}
		count++
	}
	return count
}
