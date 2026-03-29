func twoSum(numbers []int, target int) []int {
	var l = 0
	var r = len(numbers)-1
	var res []int
	for l<r{
		sum:= numbers[l]+numbers[r]
		if sum == target{
			return []int{l+1,r+1}
		}else if sum<target{
			l++
		}else{
			r--
		}
	}
	return res
}
