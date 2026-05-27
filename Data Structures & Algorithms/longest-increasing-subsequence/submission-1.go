func lengthOfLIS(nums []int) int {
	var memo = make([][]int,len(nums))
	for i := range nums{
		memo[i] = make([]int,len(nums))
		for j:= range memo[i]{
			memo[i][j] = -1
		}
	}
	return lisRecursive(nums, 0, -1, memo)
}

func lisRecursive(nums []int, curr int, prevIdx int, memo [][]int)int{
	var include, skip int
	if len(nums)== curr{
		return 0
	}
	prevMemo := prevIdx+1
	if memo[curr][prevMemo] != -1{
		return memo[curr][prevMemo]
	}
	// skip 
	skip = lisRecursive(nums, curr+1, prevIdx, memo)

	// include
	if prevIdx == -1 || nums[curr]> nums[prevIdx]{
		include = 1+lisRecursive(nums, curr+1, curr, memo)
	}
	if skip>include{
		memo[curr][prevMemo] = skip
	}else{
		memo[curr][prevMemo] = include
	}
	return memo[curr][prevMemo]
}