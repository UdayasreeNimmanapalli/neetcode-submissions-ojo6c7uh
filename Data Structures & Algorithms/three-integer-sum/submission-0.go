func threeSum(nums []int) [][]int {
	var res = make([][]int,0)
	sort.Ints(nums)
	for i:=0;i<len(nums)-2;i++{
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		left:= i+1
		right:= len(nums)-1
		for left<right{
			sum:=nums[i]+nums[left]+nums[right]
			if sum == 0{
				res = append(res,[]int{nums[i],nums[left],nums[right]})
				left++
				right--
				for left<right && nums[left-1]==nums[left]{
					left++
				}
				for left<right && nums[right+1]==nums[right]{
					right--
				}
			}else if sum > 0 {
				right--

			}else{
				left++
			}
		}
	}
	return res
}
