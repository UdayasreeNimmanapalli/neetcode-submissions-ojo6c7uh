/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
        return nil
    }
   var cloned = make(map[*Node]*Node)
   var dfs func(n *Node)*Node
   dfs = func(n *Node)*Node{
	if clone, ok:= cloned[n];ok{
		return clone
	}
	newNode := &Node{Val:n.Val}
	cloned[n] = newNode
	for _, neighbor := range n.Neighbors{
		newNode.Neighbors = append(newNode.Neighbors,dfs(neighbor))
	}
	return newNode
   }
   return dfs(node)
}
