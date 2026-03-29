func hasDuplicate(nums []int) bool {
	var dup = make(map[int]bool)
    for _,num := range nums{
		if _,ok:= dup[num];ok{
			return true
		}else{
			dup[num] = true
		}
	}
	return false
}
