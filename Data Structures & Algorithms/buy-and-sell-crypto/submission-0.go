func maxProfit(prices []int) int {
    var buy = prices[0]
    var profit = 0
    for i:=1 ; i<len(prices);i++{
        if buy>prices[i]{
            buy = prices[i]
        }else{
            profit = max(profit, prices[i]-buy)
        }
    }
    return profit
}

func max(a, b int)int{
    if a>b{
        return a
    }
    return b
}