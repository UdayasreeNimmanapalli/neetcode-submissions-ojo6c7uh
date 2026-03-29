func maxSlidingWindow(nums []int, k int) []int {
    var res = make([]int,0)
	var deque = make([]int,0)
	for i:=0;i<len(nums);i++{
		for len(deque)>0 && nums[i]>=nums[deque[len(deque)-1]]{
			deque = deque[:len(deque)-1]
		}
		deque = append(deque,i)
		if len(deque)>0 && deque[0]<=i-k{
			deque = deque[1:]
		}
		if i>=k-1{
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
