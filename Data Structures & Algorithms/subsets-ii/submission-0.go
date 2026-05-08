func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var tempArr []int
	var backtrack func(index int)
	backtrack = func(index int){
		if index == len(nums){
			res = append(res, append([]int{}, tempArr...))
			return
		}
		tempArr = append(tempArr, nums[index])
		backtrack(index+1)

		tempArr = tempArr[:len(tempArr)-1]
		for index+1<len(nums)&&nums[index+1]==nums[index]{
			index++
		}
		backtrack(index+1)
	}
	backtrack(0)
	return res
}
