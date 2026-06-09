func change(amount int, coins []int) int {
	var dp = make(map[[2]int]int)
    var dfs func(index int, amount int, dp map[[2]int]int)int
	dfs = func(index int, amount int, dp map[[2]int]int)int{
		if amount == 0{
			return 1
		}
		if index >=len(coins){
			return 0
		}
		if amount < 0{
			return 0
		}
		if val,ok:= dp[[2]int{index, amount}];ok{
			return val
		}
		pick:= dfs(index, amount-coins[index], dp)
		noPick := dfs(index+1, amount, dp)
		dp[[2]int{index, amount}] = pick+noPick
		return dp[[2]int{index, amount}]
	}
	return dfs(0, amount,dp)
}
