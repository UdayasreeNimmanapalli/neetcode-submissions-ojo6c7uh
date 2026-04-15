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
	start := 0
	end := len(lists)-1
    return mergeKListsHelper(lists, start, end)
}

func mergeKListsHelper(lists []*ListNode, start, end int)*ListNode{
	if start == end{
		return lists[start]
	}
	if start+1 == end{
		return merge2Lists(lists[start], lists[end])
	}
	mid := (start+end)/2
	left := mergeKListsHelper(lists, start, mid)
	right := mergeKListsHelper(lists, mid+1, end)
	return merge2Lists(left, right)
}
func merge2Lists(list1, list2 *ListNode)*ListNode{
	var dummy = &ListNode{}
	res := dummy
	for list1 != nil && list2 != nil{
		if list1.Val < list2.Val{
			res.Next = list1
			res = res.Next
			list1 = list1.Next
		}else{
			res.Next = list2
			res = res.Next
			list2 = list2.Next
		}
	}
	if list1!=nil{
		res.Next = list1
	}else{
		res.Next = list2
	}
	return dummy.Next
}