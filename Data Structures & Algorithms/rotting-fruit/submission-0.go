func orangesRotting(grid [][]int) int {
    var minutes = 0
	var queue = make([][2]int,0)
	var directions = [][2]int{{1,0},{0,1},{-1,0},{0,-1}}
	var freshfruits = 0
	m:= len(grid)
	n:=len(grid[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid[i][j] == 1{
				freshfruits++
			}else if grid[i][j] == 2{
				queue = append(queue, [2]int{i,j})
			}
		}
	}

	if freshfruits == 0{
		return 0
	}
	for len(queue)>0 && freshfruits>0{
		qlen:= len(queue)
		for i:=0;i<qlen;i++{
			top:=queue[0]
			queue = queue[1:]
			for _, dir := range directions{
				nr := top[0]+dir[0]
				nc := top[1]+dir[1]
				if nr>=0 && nc>=0 && nr<m && nc<n && grid[nr][nc]==1{
					freshfruits--
					grid[nr][nc]=2
					queue = append(queue, [2]int{nr,nc})
				}
			}
		}
		minutes++
	}
	if freshfruits>0{
		return -1
	}
	
	return minutes
}
