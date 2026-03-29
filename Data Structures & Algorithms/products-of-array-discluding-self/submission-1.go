func productExceptSelf(nums []int) []int {
	var zeroCount int
	var res = make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		if nums[i]==0{
			zeroCount++
		}
	}
	if zeroCount>1{
		return res
	}
	res[0] = 1
	for i:=1;i<len(nums);i++{
		res[i]= res[i-1]*nums[i-1]
	}
	n:= len(nums)
	suffix := 1
	for i:=n-1;i>=0;i--{
		res[i] *= suffix
		suffix *= nums[i]
	}
	return res
}
