func pacificAtlantic(heights [][]int) [][]int {
    var pacific = make(map[[2]int]bool)
	var atlantic = make(map[[2]int]bool)
	m:= len(heights)
	n:=len(heights[0])
	for col:=0;col<n;col++{
		//pacific
		dfs(0, col, m, n, pacific, heights, heights[0][col])
		dfs(m-1,col,m,n,atlantic, heights,heights[m-1][col])
	}
	for row:=0; row<m;row++{
		dfs(row, 0, m, n, pacific, heights, heights[row][0])
		dfs(row,n-1,m,n,atlantic, heights, heights[row][n-1])
	}
	var result [][]int
	for i:= 0;i<m;i++{
		for j:=0;j<n;j++{
			if pacific[[2]int{i,j}] && atlantic[[2]int{i,j}]{
				result = append(result,[]int{i,j})
			}
		}
	}
	return result
}

func dfs(row, col, m , n int, visited map[[2]int]bool, heights [][]int, prevHeight int){
	if row<0 || col<0 || row>=m || col>=n || visited[[2]int{row,col}] || prevHeight>heights[row][col]{
		return
	}
	visited[[2]int{row,col}]=true
	directions:= [4][2]int{{0,1},{1,0},{-1,0},{0,-1}}
	for _, dir := range directions{
		nr:= row+dir[0]
		nc:= col+dir[1]
		dfs(nr,nc,m,n,visited, heights,heights[row][col])
	}
}
