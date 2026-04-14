/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	mid := findMid(head)
	head1 := head
	secondHalf := mid.Next
	mid.Next = nil
	head2 := reverseLL(secondHalf)
	for head1 != nil && head2 != nil{
		head1Next := head1.Next
		head2Next := head2.Next
		head1.Next = head2
		head2.Next = head1Next
		head1 = head1Next
		head2 = head2Next
	}
}

func findMid(head *ListNode)*ListNode{
	slow := head
	fast := head
	for fast.Next!=nil && fast.Next.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseLL(head *ListNode)*ListNode{
	cur := head
	var prev *ListNode
	for cur!=nil{
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}