func maxArea(heights []int) int {
	var maxArea = 0
	l:=0
	r:=len(heights)-1
	for l<r{
		area := min(heights[l],heights[r])*(r-l)
		if heights[l]<heights[r]{
			l++
		}else{
			r--
		}
		maxArea = max(area, maxArea)
	}
	return maxArea
}
