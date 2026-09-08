/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    var preIdx = 0
    var indices = make(map[int]int)
    for i:=0;i<len(inorder);i++{
      indices[inorder[i]]=i
    }

    var dfs func(int, int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        if left >right {
            return nil
        }
        rootVal := preorder[preIdx]
        preIdx++
        root := &TreeNode{Val: rootVal}
        mid := indices[rootVal]
        root.Left = dfs(left, mid-1)
        root.Right = dfs(mid+1,right)

        return root
    }

    return dfs(0, len(inorder)-1)
}