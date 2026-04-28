type KthLargest struct {
    vals minHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
    kl := KthLargest{vals: minHeap{}, k:k}
	for _, num:= range nums{
		heap.Push(&kl.vals, num)
	}

	for kl.vals.Len()>kl.k{
		heap.Pop(&kl.vals)
	}
	return kl
}


func (this *KthLargest) Add(val int) int {
  heap.Push(&this.vals, val)
  for this.vals.Len()>this.k{
	heap.Pop(&this.vals)
  }
  return this.vals[0]
}

type minHeap []int

func(h minHeap)Len()int{
	return len(h)
}

func(h minHeap)Less(i, j int)bool{
	return h[i]<h[j]
}

func(h minHeap)Swap(i, j int){
	h[i],h[j]= h[j],h[i]
}

func(h *minHeap)Push(x interface{}){
	*h = append(*h, x.(int))
}

func(h *minHeap)Pop()interface{}{
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}