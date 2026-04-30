func kClosest(points [][]int, k int) [][]int {
	var h  = &p{}
	heap.Init(h)
	for _, val := range points{
		dist := (val[0]*val[0])+(val[1]*val[1])
		heap.Push(h, Pair{maxHeap: val, dist:dist})
	}
	for h.Len()>k{
		heap.Pop(h)
	}
	var res = make([][]int,0)
	for h.Len()>0{
		val := heap.Pop(h).(Pair)
		res = append(res, val.maxHeap)
	}
	return res
}

type Pair struct{
	maxHeap []int
	dist int
}
type p []Pair

func(m p)Len()int{
	return len(m)
}

func(m p)Less(i, j int)bool{
	return m[i].dist>m[j].dist
}

func(m p)Swap(i, j int){
	m[i], m[j] = m[j], m[i]
}

func(m *p)Push(x interface{}){
	*m = append(*m, x.(Pair))
}

func(m *p)Pop()interface{}{
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}