func permute(nums []int) [][]int {
	var res [][]int
	var tempArr []int
	var hMap = make(map[int]bool)
	var backtrack func()
	backtrack = func(){
		if len(tempArr)==len(nums){
			res = append(res, append([]int{}, tempArr...))
			return
		}

		for _,num:=range nums{
			if !hMap[num]{
				tempArr = append(tempArr, num)
				hMap[num]=true
				backtrack()
				tempArr = tempArr[:len(tempArr)-1]
				hMap[num]=false
			}
		}
	}
	backtrack()
	return res
}
