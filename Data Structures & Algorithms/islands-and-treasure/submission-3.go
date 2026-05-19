func islandsAndTreasure(grid [][]int) {
    m:= len(grid)
	n:= len(grid[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid[i][j]==0{
				dfs(i,j,m,n,grid,0)
			}
		}
	}
}

func dfs(row , col, m, n int, grid [][]int, dist int){
	if row<0 || col <0 || row>=m || col>=n || grid[row][col]==-1{
		return
	}

	if dist> grid[row][col]{
		return
	}
	grid[row][col] = dist

	dfs(row+1, col, m, n, grid, dist+1)
	dfs(row, col+1, m, n, grid, dist+1)
	dfs(row-1, col, m, n, grid, dist+1)
	dfs(row, col-1, m, n, grid, dist+1)
}