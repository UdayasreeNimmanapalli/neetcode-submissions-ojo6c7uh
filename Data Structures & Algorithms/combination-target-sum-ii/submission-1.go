func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var tempArr []int
	sort.Ints(candidates)
	var backtrack func(index int, target int)
	backtrack = func(index int, target int){
		if target == 0{
			res = append(res, append([]int{}, tempArr...))
			return
		}
		if len(candidates)==index ||  target<0{
			return
		}
		tempArr = append(tempArr, candidates[index])
		backtrack(index+1, target-candidates[index])
		tempArr = tempArr[:len(tempArr)-1]
		for index+1<len(candidates) && candidates[index+1]== candidates[index]{
			index++
		}
		backtrack(index+1, target)
	}
	backtrack(0, target)
	return res
}
