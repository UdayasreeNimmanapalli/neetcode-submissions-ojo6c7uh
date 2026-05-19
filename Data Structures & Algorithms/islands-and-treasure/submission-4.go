func islandsAndTreasure(grid [][]int) {
    var m = len(grid)
	var n = len(grid[0])
	var queue = make([][2]int,0)
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid[i][j]==0{
				queue = append(queue, [2]int{i,j})
			}
		}
	}
	dist:=0
	directions := [4][2]int{{1,0},{-1,0},{0,1},{0,-1}}
	for len(queue)>0{
		qlen := len(queue)
		dist++
		for i:=0;i<qlen;i++{
			top := queue[0]
			queue = queue[1:]
			for _,dir := range directions{
				nr := dir[0]+top[0]
				nc := dir[1]+top[1]
				if nr>=0 && nc>=0 && nr<m && nc<n && grid[nr][nc]==2147483647{
					grid[nr][nc]= dist
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}
}
