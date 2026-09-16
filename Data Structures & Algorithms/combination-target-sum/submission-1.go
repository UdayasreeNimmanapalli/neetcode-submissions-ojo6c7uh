func combinationSum(nums []int, target int) [][]int {
    var set = make([]int,0)
	var res = make([][]int,0)
	var dfs func(index int, target int)
	dfs = func(index int, target int){
		if index == len(nums){
			return
		}
		if target < 0 {
			return
		}

		if target == 0{
			temp := make([]int, len(set))
			copy(temp, set)
			res = append(res, temp)
			return
		}

		set = append(set, nums[index])
		dfs(index, target-nums[index])
		set = set[:len(set)-1]
		dfs(index+1, target)
	}
	dfs(0, target)
	return res
}
