func twoSum(nums []int, target int) []int {
    var sMap = make(map[int]int)
	for i, num := range nums{
		sMap[num]=i
	}
	for i,num:= range nums{
		if index,ok:= sMap[target-num];ok && index != i{
			return []int{i, index}
		}
	}
	return []int{}
}
