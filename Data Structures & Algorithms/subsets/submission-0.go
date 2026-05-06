func subsets(nums []int) [][]int {
	// use include exclude pattern
	var res = make([][]int,0)
	var currSubset = make([]int,0)
	var backtrack func(index int)
	backtrack = func(index int){
		if index == len(nums){
			res = append(res, append([]int{},currSubset...))
			return
		}

		// include index element
		currSubset = append(currSubset, nums[index])
		backtrack(index+1)
		currSubset = currSubset[:len(currSubset)-1]
		backtrack(index+1)
	}
	backtrack(0)
	return res
}
