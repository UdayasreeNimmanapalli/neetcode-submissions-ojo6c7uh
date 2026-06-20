func majorityElement(nums []int) int {
    var hMap = make(map[int]int)
	for _, num:= range nums{
		hMap[num]++
	}

	for num, count:= range hMap{
		if count>len(nums)/2{
			return num
		}
	}
	return 0
}
