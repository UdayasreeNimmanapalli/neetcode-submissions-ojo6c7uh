func lastStoneWeight(stones []int) int {
	var mHeap maxHeap
	var x, y int
	heap.Init(&mHeap)
	for _, s:= range stones{
		heap.Push(&mHeap, s)
	} 
	for mHeap.Len()>1{
		x= heap.Pop(&mHeap).(int)
		y= heap.Pop(&mHeap).(int)
		if x!=y{
			heap.Push(&mHeap, x-y)
		}
	}
	if mHeap.Len()==1{
		return mHeap[0]
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