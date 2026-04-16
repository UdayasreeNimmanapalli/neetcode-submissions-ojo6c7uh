/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
   cur := head
   for i:=0;i<k;i++{
	if cur == nil{
		return head
	}
	cur = cur.Next
   }
   resNode :=reverseNode(head,k)
   head.Next = reverseKGroup(cur,k)
   return resNode
}

func reverseNode(node *ListNode,k int)*ListNode{
	var prev *ListNode
	cur := node
	for i:=0;i<k;i++{
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}