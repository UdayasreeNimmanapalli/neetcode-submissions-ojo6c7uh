func maxProfit(prices []int) int {
    var dp = make(map[[2]int]int)
	var dfs func(index int, buying bool)int
	dfs = func(index int, buying bool)int{
		if index>=len(prices){
			return 0
		}
		if val,ok:=dp[[2]int{index, boolToInt(buying)}];ok{
			return val
		}
		if buying{
			buy := dfs(index+1, false)-prices[index]
			cooldown:= dfs(index+1, buying)
			dp[[2]int{index, boolToInt(buying)}] = max(buy, cooldown)
		}else{
			sell := dfs(index+2, true)+prices[index]
			cooldown := dfs(index+1, buying)
			dp[[2]int{index, boolToInt(buying)}] = max(sell, cooldown)
		}
		return dp[[2]int{index, boolToInt(buying)}]
	}
	return dfs(0,true)
}

func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}