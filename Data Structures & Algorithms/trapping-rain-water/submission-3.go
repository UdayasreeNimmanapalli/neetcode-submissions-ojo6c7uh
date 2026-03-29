func trap(height []int) int {
	var res = 0
	var leftMax = height[0]
	var rightMax = height[len(height)-1]
	var left = 0
	var right = len(height)-1
	for left<right{
		if leftMax<rightMax{
			left++
			leftMax = max(leftMax,height[left])
			res += leftMax - height[left]
		}else{
			right--
			rightMax = max(rightMax, height[right])
			res += rightMax - height[right]
		}
	}
	return res
}
