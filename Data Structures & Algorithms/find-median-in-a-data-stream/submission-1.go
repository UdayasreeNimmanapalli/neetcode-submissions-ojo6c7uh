type MedianFinder struct {
    arr []int
	res []int
}

var m = &minHeap{}
func Constructor() MedianFinder {
    return MedianFinder{
		arr:make([]int,0),
		res:make([]int,0),
	}
}


func (this *MedianFinder) AddNum(num int)  {
	this.res = make([]int, 0)
	var m = &minHeap{}
    this.arr = append(this.arr, num)
	heap.Init(m)
	for _,num:= range this.arr{
		heap.Push(m, num)
	}
	for m.Len()>0{
		this.res = append(this.res,heap.Pop(m).(int))
	}
	for i, j := 0, len(this.res)-1; i<j ; i, j= i+1, j-1{
		this.res[i],this.res[j]=this.res[j],this.res[i]
	}
}


func (this *MedianFinder) FindMedian() float64 {
    if len(this.res)%2!=0{
		mid := len(this.res)/2
		return float64(this.res[mid])
	}else{
		mid := len(this.res)/2
		val := (float64(this.res[mid-1])+float64(this.res[mid]))/2
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