/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
var hMap = make(map[*Node]*Node)
func copyRandomList(head *Node) *Node {
	if head == nil{
		return nil
	}
	// map[original]copy
	if val,ok:= hMap[head];ok{
		return val
	}
	copy := &Node{Val: head.Val}
	hMap[head]=copy
	copy.Next = copyRandomList(head.Next)
	copy.Random = hMap[head.Random]
	return copy
}
