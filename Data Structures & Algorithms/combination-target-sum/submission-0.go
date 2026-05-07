func combinationSum(nums []int, target int) [][]int {
    var res [][]int
	var tempArr []int

	var backtrack func(index int, target int)
	backtrack = func(index int, target int){
		if index==len(nums){
			return
		}

		if target<0{
			return
		}

		if target == 0{
			res = append(res, append([]int{}, tempArr...))
			return
		}

		tempArr = append(tempArr, nums[index])
		// include
		backtrack(index, target-nums[index])

		tempArr = tempArr[:len(tempArr)-1]
		//exclude
		backtrack(index+1, target)
	}
	backtrack(0, target)

	return res
}
