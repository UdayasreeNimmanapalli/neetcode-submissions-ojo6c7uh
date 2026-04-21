/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return nil
	}
	var res = make([]int,0)
	var q []*TreeNode
	q = append(q, root)
	for len(q)>0{
		qlen:= len(q)
		temp := []int{}
		for i:=0;i<qlen;i++{
			top := q[0]
			q = q[1:]
			if top.Left!=nil{
				q = append(q, top.Left)
			}
			if top.Right !=nil{
				q = append(q, top.Right)
			}
			temp = append(temp,top.Val)
		}
		res = append(res, temp[len(temp)-1])
	}
	return res
}
