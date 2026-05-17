func maxAreaOfIsland(grid [][]int) int {
    var visited = make(map[[2]int]bool)
    var maxArea = 0
    m:= len(grid)
    n:=len(grid[0])
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if grid[i][j]==1 && !visited[[2]int{i,j}]{
                var counter = 0
                dfs(i,j,m,n,&counter,grid,visited)
                if counter > maxArea{
                    maxArea = counter
                }
            }
        }
    }
    return maxArea
}

func dfs(row int, col int, m int, n int, counter *int, grid [][]int, visited map[[2]int]bool){
    if row<0 || col<0 || row>=m || col>=n || grid[row][col]!=1 || visited[[2]int{row, col}]{
        return
    }

    visited[[2]int{row,col}]=true
    *counter++
    dfs(row+1, col,m,n,counter, grid, visited)
    dfs(row, col+1,m,n,counter, grid, visited)
    dfs(row-1, col,m,n,counter, grid, visited)
    dfs(row, col-1,m,n,counter, grid, visited)
}