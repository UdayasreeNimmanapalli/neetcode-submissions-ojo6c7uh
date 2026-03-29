func longestConsecutive(nums []int) int {
	 var maxCount = 0
	 var hMap = make(map[int]struct{})
	 for _, num:= range nums{
		hMap[num]= struct{}{}
	 }
	 for num:= range hMap{
		if _,ok:= hMap[num-1];!ok{
			count :=1
			for{
				if _,ok:=hMap[num+count];ok{
				count++
				}else{
					break
				} 
			}
			if count>maxCount{
				maxCount = count
			}
		}
	 }
	 return maxCount
}
