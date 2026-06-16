/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	var pq = &minHeap{}
	var res = &ListNode{}
	var dummy = res
	heap.Init(pq)
	for i:=0;i<len(lists);i++{
		heap.Push(pq,lists[i])
	}

	for pq.Len()>0{
		node:= heap.Pop(pq).(*ListNode)
		dummy.Next = node
		dummy = dummy.Next

		if node.Next!=nil{
			node = node.Next
			heap.Push(pq, node)
		}
	}
	return res.Next
}

type minHeap []*ListNode

func(p *minHeap)Len()int{
	return len(*p)
}

func(p *minHeap)Less(i, j int)bool{
	return (*p)[i].Val<(*p)[j].Val
}

func(p *minHeap)Swap(i, j int){
	 (*p)[i],(*p)[j]=(*p)[j],(*p)[i]
}

func(p *minHeap)Push(x interface{}){
	*p = append(*p, x.(*ListNode))
}

func(p *minHeap)Pop()interface{}{
	top := (*p)[len(*p)-1]
	(*p) = (*p)[:len(*p)-1]
	return top
}