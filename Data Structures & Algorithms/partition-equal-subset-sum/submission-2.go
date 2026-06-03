func canPartition(nums []int) bool {
    sum := 0
	for _,num:= range nums{
		sum += num
	}
	if sum%2!=0{
		return false
	}
	target := sum/2
	var dp = make([]bool,target+1)
	dp[0]=true
	for i:=0;i<len(nums);i++{
		for j:=target;j>=nums[i];j--{
			if dp[j-nums[i]]{
				dp[j]=true
			}
		}
	}
	return dp[target]
}