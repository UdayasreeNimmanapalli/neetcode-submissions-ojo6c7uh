func findTargetSumWays(nums []int, target int) int {
    var dfs func(index int, total int)int
	dfs = func(index int, total int)int{
		if index==len(nums){
			if total == target{
				return 1
			}
			return 0
		}

		return dfs(index+1, total+nums[index])+dfs(index+1,total-nums[index])
	}
	return dfs(0, 0)
}
