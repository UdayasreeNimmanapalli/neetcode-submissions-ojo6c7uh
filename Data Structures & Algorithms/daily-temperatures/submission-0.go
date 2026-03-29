func dailyTemperatures(temperatures []int) []int {
	var res = make([]int,len(temperatures))
	var stack = make([]int,0)
	for i := range temperatures{
		for len(stack)>0 && temperatures[i]> temperatures[stack[len(stack)-1]]{
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[topIdx] = i-topIdx
		}
		stack = append(stack, i)
	}
	return res
}
