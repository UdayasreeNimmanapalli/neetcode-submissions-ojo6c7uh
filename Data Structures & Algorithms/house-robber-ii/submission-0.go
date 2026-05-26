func rob(nums []int) int {
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
    return max(robLinear(nums[0:len(nums)-1]),robLinear(nums[1:len(nums)]))
}

func robLinear(nums []int)int{
	var memo = make(map[int]int)
	var dfs func(index int)int
	dfs = func(index int)int{
		if index>=len(nums){
			return 0
		}
		if _,ok:= memo[index];ok{
			return memo[index]
		}
		memo[index] = max(nums[index]+dfs(index+2), dfs(index+1))
		return memo[index]
	}
	return dfs(0)
}