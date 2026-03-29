func maxProfit(prices []int) int {
	var buy = prices[0]
	var maxProfit = 0
	for i:=1;i<len(prices);i++{
		if prices[i]<buy{
			buy = prices[i]
		}else{
			profit := prices[i]-buy
			if profit>maxProfit{
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
