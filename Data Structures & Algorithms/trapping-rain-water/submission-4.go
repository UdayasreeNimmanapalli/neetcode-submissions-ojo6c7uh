func trap(height []int) int {
	var water = 0
	var maxLeft = make([]int,len(height))
	var maxRight = make([]int, len(height))
	maxLeft[0]=0
	for i:=1;i<len(height);i++{
		maxLeft[i] = max(maxLeft[i-1],height[i-1])
	}
	n:= len(height)
	maxRight[n-1]=0
	for i:=n-2;i>=0;i--{
		maxRight[i] = max(maxRight[i+1],height[i+1])
	}
	for i:=0;i<len(height);i++{
		h := min(maxLeft[i],maxRight[i])-height[i]
		if h<0{
			h=0
		}
		water += h
	}
	return water
}
