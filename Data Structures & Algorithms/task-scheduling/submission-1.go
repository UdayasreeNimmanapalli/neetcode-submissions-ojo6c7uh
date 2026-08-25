func leastInterval(tasks []byte, n int) int {
	var mHeap = &maxHeap{}
	heap.Init(mHeap)
	var queue = make([][2]int,0)
	var freq = make(map[byte]int)
	for _,task:= range tasks{
		freq[task]++
	}
	for _,count:= range freq{
		heap.Push(mHeap, count)
	}
	var time = 0
	for len(queue)>0 || mHeap.Len()>0{
		time++
		if mHeap.Len() == 0{
			time = queue[0][1]
		}else{
			val:= heap.Pop(mHeap).(int)-1
			if val>0{
				queue = append(queue, [2]int{val, time+n})
			}
		}

		if len(queue)>0 && time == queue[0][1]{
			heap.Push(mHeap, queue[0][0])
			queue = queue[1:]
		}
	}
	return time
}


type maxHeap []int

func(m maxHeap)Len()int{
	return len(m)
}

func(m maxHeap)Less(i, j int)bool{
	return m[i]>m[j]
}

func(m maxHeap)Swap(i, j int){
	m[i],m[j]=m[j],m[i]
}

func(m *maxHeap)Push(x interface{}){
	*m = append(*m, x.(int))
}

func(m *maxHeap)Pop()interface{}{
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}