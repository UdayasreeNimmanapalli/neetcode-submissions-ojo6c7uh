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
    cur := head
	count := 0
	for cur != nil{
		cur = cur.Next
		count++
	}
	if n==count{
		return head.Next
	}
	newCount := count-n-1
	dummy := head
	for newCount!=0 && dummy!=nil{
		dummy = dummy.Next
		newCount--
	}
	if dummy.Next!=nil{
		dummy.Next = dummy.Next.Next
	}
	return head
}
