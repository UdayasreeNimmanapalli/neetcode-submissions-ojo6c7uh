func missingNumber(nums []int) int {
	sum:=0
	sum1 :=0
	for i:=0;i<=len(nums);i++{
		sum += i
	}
	 for _, val := range nums{
		sum1+=val
	 }
	 return sum-sum1
}
