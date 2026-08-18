type KthLargest struct {
    k int
	nums []int
}


func Constructor(k int, nums []int) KthLargest {
   kl := KthLargest{k:k, nums:nums}
   heap.Init(&kl)
   for len(kl.nums)>k{
	 heap.Pop(&kl)
   }
   return kl
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this, val)
	if len(this.nums)>this.k{
		heap.Pop(this)
	}
	return this.nums[0]
}


func(m *KthLargest)Len()int{
	return len(m.nums)
}

func(m *KthLargest)Less(i, j int)bool{
	return m.nums[i]<m.nums[j]
}

func(m *KthLargest)Swap(i, j int){
	m.nums[i],m.nums[j] = m.nums[j],m.nums[i]
}

func(m *KthLargest)Push(x interface{}){
	(m.nums) = append((m.nums), x.(int))
}

func(m *KthLargest)Pop()interface{}{
	top:=(m.nums)[len(m.nums)-1]
	(m.nums) = (m.nums)[:len(m.nums)-1]
	return top
}