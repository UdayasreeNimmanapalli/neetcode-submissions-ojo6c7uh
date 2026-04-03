func carFleet(target int, position []int, speed []int) int {
	var stack = make([]float64,0)
	var sorted = make([][2]int,len(position))
	for i:= range position{
		sorted[i]=[2]int{position[i], speed[i]}
	}
	sort.Slice(sorted,func(i,j int)bool{
		return sorted[i][0]>sorted[j][0]
	})
	for _,pair:=range sorted{
		distToTarget := float64(target-pair[0])/float64(pair[1])
		stack = append(stack,distToTarget)
		if len(stack) >= 2 && (stack[len(stack)-1]<=stack[len(stack)-2]){
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}
