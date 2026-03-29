func maxArea(heights []int) int {
	var left = 0
	var right = len(heights)-1
	var maxArea = math.MinInt
	for left<right{
		area := min(heights[left],heights[right]) * (right-left)
		if area > maxArea{
			maxArea = area
		}
		if heights[left]<heights[right]{
			left++
		}else {
			right--
		}
	}
	return maxArea
}
