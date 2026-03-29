func maxProfit(prices []int) int {
	var maxProfit = math.MinInt
	var minValue = prices[0]
	for i:=1;i<len(prices);i++{
		if minValue>prices[i]{
			minValue = prices[i]
		}else{
			maxProfit = max(maxProfit, (prices[i]-minValue))
		}
	}
	if maxProfit<0{
		maxProfit = 0
	}
	return maxProfit
}
