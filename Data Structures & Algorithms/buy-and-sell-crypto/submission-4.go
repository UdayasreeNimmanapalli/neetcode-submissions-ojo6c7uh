func maxProfit(prices []int) int {
	buy:=prices[0]
	var maxProfit = 0
	for i:=1;i<len(prices);i++{
		if buy>prices[i]{
			buy = prices[i]
		}else{
			maxProfit = max(maxProfit, prices[i]-buy)
		}
	}
	return maxProfit
}
