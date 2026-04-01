func dailyTemperatures(temperatures []int) []int {
	var res = make([]int,len(temperatures))
	// for i:=0;i<len(temperatures);i++{
	// 	temperatures[i] = 0
	// }
	var stack = make([]int,0)
	for i:=0;i<len(temperatures);i++{
		for len(stack)>0 && temperatures[i]>temperatures[stack[len(stack)-1]]{
			res[stack[len(stack)-1]] = i-stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} 
		stack = append(stack, i)
	}
	return res
}
