func findKthLargest(nums []int, k int) int {
	var mHeap = &minHeap{}
	heap.Init(mHeap)
	for _, num:= range nums{
		heap.Push(mHeap, num)
	}
	for mHeap.Len()>k{
		heap.Pop(mHeap)
	}
	return heap.Pop(mHeap).(int)
}

type minHeap []int

func(m minHeap)Len()int{
	return len(m)
}

func(m minHeap)Less(i, j int)bool{
	return m[i]<m[j]
}

func(m minHeap)Swap(i, j int){
	m[i], m[j] = m[j], m[i]
}

func(m *minHeap)Push(x interface{}){
	*m = append(*m, x.(int))
}

func(m *minHeap)Pop()interface{}{
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}