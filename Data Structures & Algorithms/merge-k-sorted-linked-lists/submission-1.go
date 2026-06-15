/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists)==0{
		return nil
	}
	for i:=1;i<len(lists);i++{
		lists[i] = merge2Lists(lists[i-1],lists[i])
	}
	return lists[len(lists)-1]
}

func merge2Lists(l1 *ListNode, l2 *ListNode)*ListNode{
	var res = &ListNode{}
	dummy := res
	for l1!=nil && l2!=nil{
		if l1.Val<=l2.Val{
			dummy.Next = l1
			dummy= dummy.Next
			l1=l1.Next
		}else{
			dummy.Next = l2
			dummy=dummy.Next
			l2 = l2.Next
		}
	}
	if l1!=nil{
		dummy.Next = l1
	}else{
		dummy.Next = l2
	}
	return res.Next
}
