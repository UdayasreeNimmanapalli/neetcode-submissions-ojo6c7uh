func longestConsecutive(nums []int) int {
	var present = make(map[int]struct{})
	for _,num := range nums{
		present[num]=struct{}{}
	}
	var res  = 0
	
	for _,num  := range nums{
		if _,ok:= present[num-1];!ok{
			length:=1
			for{
				if _,ok:= present[num+length];ok{
					length++
				}else{
					break
				}
			}
			if length>res{
				res = length
			}	
		}
	}
	return res
}
