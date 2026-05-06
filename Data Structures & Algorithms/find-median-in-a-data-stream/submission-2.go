type MedianFinder struct {
    small *maxHeap
	large *minHeap
}

func Constructor() MedianFinder {
    return MedianFinder{
		small:&maxHeap{},
		large:&minHeap{},
	}
}


func (this *MedianFinder) AddNum(num int)  {
   heap.Push(this.small, num)
	if this.small.Len()>0 && this.large.Len()>0 && (*this.small)[0] > (*this.large)[0]{
		val := heap.Pop(this.small).(int)
		heap.Push(this.large, val)
	}

	if this.small.Len()>=this.large.Len()+1{
		val:=heap.Pop(this.small).(int)
		heap.Push(this.large, val)
	}

	if this.large.Len()>this.small.Len(){
		val:= heap.Pop(this.large).(int)
		heap.Push(this.small, val)
	}
}


func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len()!= this.large.Len(){
		return float64((*this.small)[0])
	}else{
		val := (float64((*this.small)[0])+float64((*this.large)[0]))/2
		return val
	}
	return 0
}

type minHeap []int

func(m minHeap)Len()int{
	return len(m)
}

func(m minHeap)Less(i, j int)bool{
	return m[i]<m[j]
}

func(m minHeap)Swap(i, j int){
	m[i],m[j] = m[j], m[i]
}

func(m *minHeap)Push(x interface{}){
	*m=append(*m, x.(int))
}

func(m *minHeap)Pop()interface{}{
	x:=(*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

type maxHeap []int

func(m maxHeap)Len()int{
	return len(m)
}

func(m maxHeap)Less(i, j int)bool{
	return m[i]>m[j]
}

func(m maxHeap)Swap(i, j int){
	m[i],m[j] = m[j], m[i]
}

func(m *maxHeap)Push(x interface{}){
	*m=append(*m, x.(int))
}

func(m *maxHeap)Pop()interface{}{
	x:=(*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}