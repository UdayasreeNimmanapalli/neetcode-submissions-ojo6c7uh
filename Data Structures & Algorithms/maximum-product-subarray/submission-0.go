func maxProduct(nums []int) int {
    // Modified Kadanes Algorithm
	var maxCur = nums[0]
	var minCur = nums[0]
	var globalMax = nums[0]

	for i:=1;i<len(nums);i++{
		num := nums[i]
		if num<0{
			maxCur, minCur = minCur, maxCur
		}
		maxCur = max(num,num*maxCur)
		minCur = min(num,num*minCur)
		if maxCur>globalMax{
			globalMax = maxCur
		}
	}
	return globalMax
}
