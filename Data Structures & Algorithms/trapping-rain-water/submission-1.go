func trap(height []int) int {
	var res = 0
	var leftMax = make([]int,len(height))
	leftMax[0]= height[0]
	var rightMax = make([]int,len(height))
	rightMax[len(height)-1]=height[len(height)-1]
	for i:=1;i<len(height);i++{
		leftMax[i]=max(leftMax[i-1],height[i])
	}
	for i:=len(height)-2;i>=0;i--{
		rightMax[i]=max(rightMax[i+1],height[i])
	}
	for i:=0;i<=len(height)-1;i++{
		level := min(leftMax[i],rightMax[i])
		res += level-height[i]
	}
	return res
}
