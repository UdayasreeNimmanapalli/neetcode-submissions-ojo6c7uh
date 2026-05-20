func solve(board [][]byte) {
    m:=len(board)
	n:=len(board[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			var regions = make([][2]int,0)
			var touchesBorder = false
			if board[i][j]=='O'{
				dfs(i,j,m,n,&regions, board, &touchesBorder)
				for _,reg := range regions{
					if touchesBorder{
						board[reg[0]][reg[1]] = 'O'
					}else{
						board[reg[0]][reg[1]] = 'X'
					}
				}
			}
		}
	}
}


func dfs(row, col , m, n int, regions *[][2]int, board [][]byte, touchesBorder *bool){
	if row<0 || col<0 || row>=m || col>=n ||board[row][col]!='O'{
		return
	}
	board[row][col]='V'
	*regions = append(*regions, [2]int{row,col})
	
	if row == 0 || row == m-1 || col == 0 || col == n-1{
		*touchesBorder=true
	}
	directions := [4][2]int{{1,0},{0,1},{-1,0},{0,-1}}
	for _, dir := range directions{
		nr:= row+dir[0]
		nc:= col+dir[1]
		dfs(nr,nc,m,n,regions, board, touchesBorder)
	}
}