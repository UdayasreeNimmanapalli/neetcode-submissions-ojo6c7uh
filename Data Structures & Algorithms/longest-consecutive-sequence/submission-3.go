func longestConsecutive(nums []int) int {
	var hMap = make(map[int]struct{},0)
	for _, val := range nums{
		hMap[val]=struct{}{}
	}
	var maxCount = 0
	for _, num:= range nums{
		count :=1
		for {
			if _, ok:= hMap[num+1];ok{
				num++
				count++
			}else{
				break	
			}
		}
		if maxCount<count{
			maxCount = count
		}
	}
	return maxCount
}
