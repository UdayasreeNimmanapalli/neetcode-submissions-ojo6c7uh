/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
      return nil
    }
    var hMap = make(map[int]int)
    root:= &TreeNode{Val:preorder[0]}
    for i:= 0;i<len(inorder);i++{
      hMap[inorder[i]]=i
    }
    mid := hMap[preorder[0]]
    root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
    return root
}
