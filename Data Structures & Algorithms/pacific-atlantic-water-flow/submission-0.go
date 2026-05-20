func pacificAtlantic(heights [][]int) [][]int {
    var result = make([][]int,0)
	m:= len(heights)
	n:=len(heights[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			var visitedA = make(map[[2]int]bool)
			var visitedP = make(map[[2]int]bool)
			visitedA[[2]int{i,j}]=true
			visitedP[[2]int{i,j}]=true
			if canReachPacific(i,j,m,n,heights, visitedA)&&
			canReachAtlantic(i,j,m,n,heights, visitedP){
				result = append(result, []int{i,j})
			}
		}
	}
	return result
}

func canReachPacific(row, col , m, n int, heights[][]int, visitedA map[[2]int]bool)bool{
	if row == 0 || col ==0{
		return true
	}
	directions := [][2]int{{1,0},{0,1},{-1,0},{0,-1}}
	for _, dir := range directions{
		nr := row+dir[0]
		nc:= col+dir[1]
		if nr<0 || nc<0 || nr>=m  || nc>=n || visitedA[[2]int{nr,nc}] || heights[row][col]<heights[nr][nc]{
			continue
		}
		visitedA[[2]int{nr,nc}]=true
		if canReachPacific(nr, nc, m, n, heights, visitedA){
			return true
		}
	}
	return false
}

func canReachAtlantic(row, col , m, n int, heights[][]int, visitedP map[[2]int]bool)bool{
	if row == m-1 || col == n-1{
		return true
	}
	directions := [][2]int{{1,0},{0,1},{-1,0},{0,-1}}
	for _, dir := range directions{
		nr := row+dir[0]
		nc:= col+dir[1]
		if nr<0 || nc<0 || nr>=m  || nc>=n || visitedP[[2]int{nr,nc}] || heights[row][col]<heights[nr][nc]{
			continue
		}
		visitedP[[2]int{nr,nc}]=true
		if canReachAtlantic(nr, nc, m, n, heights, visitedP){
			return true
		}
	}
	return false
}