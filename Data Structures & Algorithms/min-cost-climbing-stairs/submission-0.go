func minCostClimbingStairs(cost []int) int {
	var memo = make(map[int]int,0)
	var dfs func(index int)int
	dfs = func(index int)int{
		if index>=len(cost){
			return 0
		}
		if _,ok:= memo[index];ok{
			return memo[index]
		}
		memo[index] = cost[index]+min(dfs(index+1),dfs(index+2))
		return memo[index]
	}
	return min(dfs(0),dfs(1))
}
