func maxArea(heights []int) int {
	l:=0
	r := len(heights)-1
	maxArea := 0
	for l<r{
		area := min(heights[l],heights[r])*(r-l)
		if area >maxArea{
			maxArea = area
		}
		if heights[l]< heights[r]{
			l++
		}else{
			r--
		}
	}
	return maxArea
}
