func trap(height []int) int {
	var water = 0
	var leftMax = 0
	var rightMax = 0
	l:=0
	r:=len(height)-1
	for l<r{
		if height[l]<=height[r]{
			if leftMax<height[l]{
				leftMax=height[l]
			}else{
				water += leftMax-height[l]
			}
			l++
		}else{
			if rightMax<height[r]{
				rightMax=height[r]
			}else{
				water += rightMax-height[r]
			}
			r--
		}
	}
	return water
}
