func search(nums []int, target int) int {
	var l = 0
	var r = len(nums)-1
	for l<=r{
		mid := (l+r)/2
		if nums[mid] == target{
			return mid
		}else if nums[mid]<target{
			l=mid+1
		}else{
			r=mid-1
		}
	}
	return -1
}
