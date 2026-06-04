func uniquePaths(m int, n int) int {
	var dp = make([][]int,m)
	for i:= range dp{
		dp[i]= make([]int,n)
	}
    return rec(0,0,m,n,dp)
}

func rec(row, col , m, n int, dp [][]int)int{
	if row == m-1 && col == n-1{
		return 1
	}
	if row>=m || col>=n{
		return 0
	}
	if dp[row][col]!=0{
		return dp[row][col]
	}
	dp[row][col] = rec(row, col+1,m,n,dp)+rec(row+1, col, m,n,dp)
	return dp[row][col]
}