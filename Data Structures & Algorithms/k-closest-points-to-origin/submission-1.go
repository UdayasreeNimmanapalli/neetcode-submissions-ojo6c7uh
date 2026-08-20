func kClosest(points [][]int, k int) [][]int {
	var mh = & maxHeap{}
	heap.Init(mh)
	for _,val:= range points{
		dist := (val[0]*val[0])+(val[1]*val[1])
		heap.Push(mh, Pair{val, dist})
	}

	for mh.Len()>k{
		heap.Pop(mh)
	}

	result := make([][]int,0)
	for mh.Len()>0{
		val := heap.Pop(mh).(Pair)
		result = append(result, val.point)
	}
	return result
}

type Pair struct{
	point []int
	dist int
}

type maxHeap []Pair

func(m maxHeap)Len()int{
	return len(m)
}

func(m maxHeap)Less(i, j int)bool{
	return (m[i]).dist>(m[j]).dist
}

func(m maxHeap)Swap(i, j int){
	m[i],m[j] = m[j],m[i]
}

func(m *maxHeap)Push(x interface{}){
	*m = append(*m, x.(Pair))
}

func(m *maxHeap)Pop()interface{}{
	top := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return top
}