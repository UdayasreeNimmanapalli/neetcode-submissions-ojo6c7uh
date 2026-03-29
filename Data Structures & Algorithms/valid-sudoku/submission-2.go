func isValidSudoku(board [][]byte) bool {
	var rowMap = make([]map[byte]bool,9)
	var colMap = make([]map[byte]bool,9)
	var squareMap = make([]map[byte]bool,9)

	for i:=0; i<9 ; i++{
		rowMap[i] = make(map[byte]bool)
		colMap[i] = make(map[byte]bool)
		squareMap[i] = make(map[byte]bool)
	}

	for i:=0; i<9;i++{
		for j:= 0;j<9;j++{
			if board[i][j]=='.'{
				continue
			}
			val := board[i][j]
			squareIndex := (i/3)*3 + (j/3)
			if rowMap[i][val] || colMap[j][val] || squareMap[squareIndex][val]{
				return false
			}
			rowMap[i][val] = true
			colMap[j][val] = true
			squareMap[squareIndex][val]= true
		}
	}
	return true
}
