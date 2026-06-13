/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil{
		return nil
	}
	dummy := &ListNode{Val:0, Next:head}
	curr:= head
	curr1:= dummy
	count:=0
	for curr!=nil{
		count++
		curr=curr.Next
	}
	newCount := count-n
	for newCount>0 && curr1!=nil{
		newCount--
		curr1 = curr1.Next
	}
	if curr1.Next!=nil{
		curr1.Next=curr1.Next.Next
	}
	return dummy.Next
}
