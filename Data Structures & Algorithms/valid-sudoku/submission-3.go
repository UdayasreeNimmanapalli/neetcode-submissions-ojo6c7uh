func isValidSudoku(board [][]byte) bool {
	var row = make([]map[byte]bool,9)
	var col = make([]map[byte]bool,9)
	var sqr = make([]map[byte]bool,9)

	for i:=0;i<9;i++{
		row[i]=make(map[byte]bool)
		col[i]=make(map[byte]bool)
		sqr[i]=make(map[byte]bool)
	}
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]=='.'{
				continue
			}
			val := board[i][j]
			sqrIdx := (i/3)*3+(j/3)
			if row[i][val] || col[j][val]|| sqr[sqrIdx][val]{
				return false
			}
			row[i][val]=true
			col[j][val]= true
			sqr[sqrIdx][val]=true
		}
	}
	return true
}
