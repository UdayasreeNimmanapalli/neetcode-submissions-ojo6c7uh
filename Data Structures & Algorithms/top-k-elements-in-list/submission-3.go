func topKFrequent(nums []int, k int) []int {
	var mHeap = &minHeap{}
	heap.Init(mHeap)
	var freq = make(map[int]int)
	for _, num := range nums{
		freq[num]++
	}
	for num, count := range freq{
		heap.Push(mHeap, pair{num, count})
	}

	for mHeap.Len()>k{
		heap.Pop(mHeap)
	}
	var res = make([]int,0)
	for _, val:= range *mHeap{
		res = append(res, val.num)
	}
	return res
}

type pair struct{
	num int
	count int
}

type minHeap []pair
func(m minHeap)Len()int{
	return len(m)
}

func(m minHeap)Less(i, j int)bool{
	return m[i].count<m[j].count
}

func(m minHeap)Swap(i, j int){
	m[i], m[j] = m[j], m[i]
}

func(m *minHeap)Push(x interface{}){
	*m = append(*m, x.(pair))
}

func(m *minHeap)Pop()interface{}{
	val:= (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}