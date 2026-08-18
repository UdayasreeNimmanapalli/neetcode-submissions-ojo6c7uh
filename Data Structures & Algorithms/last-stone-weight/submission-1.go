func lastStoneWeight(stones []int) int {
	var mh = &maxHeap{}
	heap.Init(mh)
	for _,val:=range stones{
		heap.Push(mh,val)
	}
	for mh.Len()>1{
		x:= heap.Pop(mh).(int)
		y:= heap.Pop(mh).(int)
		if x!=y{
			heap.Push(mh, x-y)
		}
	}
	if mh.Len()==1{
		return (*mh)[0]
	}
	return 0
}

type maxHeap []int

func(h maxHeap)Len()int{
	return len(h)
}

func(h maxHeap)Less(i, j int)bool{
	return h[i]>h[j]
}

func(h maxHeap)Swap(i, j int){
	h[i],h[j]=h[j], h[i]
}

func(h *maxHeap)Push(x interface{}){
	*h = append(*h, x.(int))
}

func(h *maxHeap)Pop()interface{}{
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}