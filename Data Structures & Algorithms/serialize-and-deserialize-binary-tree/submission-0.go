/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var res = make([]string,0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode){
		if node == nil{
			res = append(res, "N")
			return
		}
		strNode:=strconv.Itoa(node.Val)
		res = append(res, strNode)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(res,",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    nodes := strings.Split(data,",")
	preIdx := 0
	var dfs func()*TreeNode
	dfs = func()*TreeNode{
		if nodes[preIdx]=="N"{
			preIdx++
			return nil
		}
		val,_ := strconv.Atoi(nodes[preIdx])
		node := &TreeNode{Val:val}
		preIdx++
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}
