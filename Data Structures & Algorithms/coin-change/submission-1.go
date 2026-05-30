func coinChange(coins []int, amount int) int {
	if len(coins)==0 || amount ==0{
		return 0
	}
	var memo = make(map[int]int)
	var dfs func(amt int, memo map[int]int)int
	dfs = func(amt int, memo map[int]int)int{
		if val,ok := memo[amt];ok{
			return val
		}
		if amt == 0{
			return 0
		}
		res := math.MaxInt32
		for _, coin := range coins{
			if amt-coin>=0{
				res = min(res, 1+dfs(amt-coin, memo))
			}
		}
		memo[amt] = res
		return res
	}
	minCoins := dfs(amount,memo)
	if minCoins >= math.MaxInt32{
		return -1
	}
	return minCoins
}
