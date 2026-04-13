/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var res = &ListNode{}
	dummy := res
	cur1:=list1
	cur2:=list2
	for cur1!=nil && cur2!=nil{
		if cur1.Val < cur2.Val{
			dummy.Next = cur1
			dummy = dummy.Next
			cur1 = cur1.Next
		}else{
			dummy.Next = cur2
			dummy = dummy.Next
			cur2=cur2.Next
		}
	}
	if cur1!=nil{
		dummy.Next = cur1
	}else {
		dummy.Next = cur2
	}
	return res.Next
}
