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
	var prefix = make([]int,len(nums))
	var suffix = make([]int,len(nums))
	prefix[0] = 1
	for i:=1;i<len(nums);i++{
		prefix[i]= prefix[i-1]*nums[i-1]
	}
	n:= len(nums)
	suffix[n-1] = 1
	for i:=n-2;i>=0;i--{
		suffix[i]= suffix[i+1]*nums[i+1]
	}

	for j:=0;j<len(nums);j++{
		res[j]= suffix[j]*prefix[j]
	}
	return res
}
