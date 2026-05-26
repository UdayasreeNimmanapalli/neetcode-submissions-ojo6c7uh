func rob(nums []int) int {
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	if len(nums)==2{
		return max(nums[0],nums[1])
	}
    return max(robLinear(nums[0:len(nums)-1]),robLinear(nums[1:len(nums)]))
}

func robLinear(nums []int)int{
	var dp = make([]int,len(nums))
	dp[0]= nums[0]
	dp[1] = max(nums[0],nums[1])
	for i:=2;i<len(nums);i++{
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}