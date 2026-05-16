func numIslands(grid [][]byte) int {
    var islands = 0
	var visited = make(map[[2]int]bool)
	m:= len(grid)
	n:= len(grid[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if !visited[[2]int{i,j}] && grid[i][j]=='1'{
				islands++
				dfs(i,j,m,n,grid,visited)
			}
		}
	}
	return islands
}


func dfs(row int, col int, m int, n int, grid [][]byte, visited map[[2]int]bool){
	if row<0 || col<0 || row>=m || col>=n || grid[row][col]!='1'|| visited[[2]int{row,col}]{
		return
	}
	visited[[2]int{row,col}] = true
	dfs(row+1, col, m, n, grid, visited)
	dfs(row, col+1, m, n, grid, visited)
	dfs(row-1, col, m, n, grid, visited)
	dfs(row, col-1, m, n, grid, visited)
}