func exist(board [][]byte, word string) bool {
	m := len(board)
	n:= len(board[0])
	index :=0 
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if word[index]==board[i][j]{
				if dfs(board, i, j, index, m, n, word){
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, row int, col int, index int, m, n int, word string)bool{
	if row<0 || col<0 || row>=m || col>=n || board[row][col]!= word[index]{
		return false
	}
	if index==len(word)-1{
		return true
	}
	temp := board[row][col]

	board[row][col] = '#'

	found := dfs(board,row+1, col, index+1,m,n, word)||
	dfs(board,row, col+1, index+1,m,n, word)||
	dfs(board,row-1, col, index+1,m,n, word)||
	dfs(board,row, col-1, index+1,m,n, word)

	board[row][col] = temp
	
	return found
}