func combinationSum2(candidates []int, target int) [][]int {
	var set = make([]int,0)
	var res = make([][]int,0)
	sort.Ints(candidates)
	var dfs func(index int, target int)
	dfs = func(index int, target int){
		if target == 0{
			temp := make([]int, len(set))
			copy(temp, set)
			res = append(res, temp)
			return
		}

		if target < 0 || index==len(candidates){
			return 
		}

		set = append(set, candidates[index])
		dfs(index+1, target-candidates[index])
		set = set[:len(set)-1]
		for index+1<len(candidates) && candidates[index]==candidates[index+1]{
			index++
		}
		dfs(index+1, target)
	}
	dfs(0, target)
	return res
}
