type MedianFinder struct {
    small *maxHeap
	large *minHeap
}


func Constructor() MedianFinder {
	small:= &maxHeap{}
	large := &minHeap{}
	heap.Init(small)
	heap.Init(large)
    return MedianFinder{
		small: small,
		large : large,
	}
}


func (this *MedianFinder) AddNum(num int)  {
   if this.large.Len()>0{
	if num > ((*this.large)[0]){
		heap.Push(this.large, num)
	}else{
		heap.Push(this.small, num)
	}
   }else{
		heap.Push(this.small, num)
   }

   if this.large.Len()>this.small.Len()+1{
	heap.Push(this.small, heap.Pop(this.large))
   }
   if this.small.Len()>this.large.Len()+1{
	heap.Push(this.large, heap.Pop(this.small))
   }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len()>this.large.Len(){
		return float64((*this.small)[0])
	}

	if this.large.Len()>this.small.Len(){
		return float64((*this.large)[0])
	}

	return (float64((*this.large)[0]+ (*this.small)[0]))/2.0
}

type minHeap []int

func(m minHeap)Len()int{
	return len(m)
}

func(m minHeap)Less(i, j int)bool{
	return m[i]< m[j]
}

func(m minHeap)Swap(i, j int){
	m[i],m[j]=m[j],m[i]
}

func(m *minHeap)Push(x interface{}){
	*m = append(*m, x.(int))
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
	return m[i]> m[j]
}

func(m maxHeap)Swap(i, j int){
	m[i],m[j]=m[j],m[i]
}

func(m *maxHeap)Push(x interface{}){
	*m = append(*m, x.(int))
}

func(m *maxHeap)Pop()interface{}{
	x:=(*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}
