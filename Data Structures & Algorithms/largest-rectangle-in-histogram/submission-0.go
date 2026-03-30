func largestRectangleArea(heights []int) int {
	var maxArea = 0
	var nearestSmallerLeft = make([]int,len(heights))
	for i:=0;i<len(heights);i++{
		nearestSmallerLeft[i] = -1
	}
	var leftStack = make([]int,0)
	for i:=0;i<len(heights);i++{
		for len(leftStack)>0 && heights[leftStack[len(leftStack)-1]]>=heights[i]{
			leftStack = leftStack[:len(leftStack)-1]
		}
		if len(leftStack)>0{
			nearestSmallerLeft[i] = leftStack[len(leftStack)-1]
		}
		leftStack = append(leftStack,i)
	}
	 var nearestSmallerRight = make([]int,len(heights))
	 n := len(heights)
	 for i:=0;i<n;i++{
		nearestSmallerRight[i]=n
	 }
	 var rightStack = make([]int,0)
	 for i:=n-1;i>=0;i--{
		for len(rightStack)>0 && heights[rightStack[len(rightStack)-1]]>=heights[i]{
			rightStack = rightStack[:len(rightStack)-1]
		}
		if len(rightStack)>0{
			nearestSmallerRight[i]=rightStack[len(rightStack)-1]
		}
		rightStack = append(rightStack, i)
	 }

	 for i:=0;i<len(heights);i++{
		area := heights[i]*(nearestSmallerRight[i]-nearestSmallerLeft[i]-1)
		if area > maxArea{
			maxArea = area
		}
	 }
	 return maxArea
}
