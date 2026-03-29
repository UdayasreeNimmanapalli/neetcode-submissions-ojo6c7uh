func isValidSudoku(board [][]byte) bool {
	for row := 0; row<9; row++{
		present := make(map[byte]bool)
		for i:=0;i<9;i++{
			if board[row][i] == '.'{
				continue
			}
			if present[board[row][i]]{
				return false
			}
			present[board[row][i]] = true
		}
	}

	for col := 0; col<9 ; col++{
		present := make(map[byte]bool)
		for j:=0;j<9;j++{
			if board[j][col]=='.'{
				continue
			}
			if present[board[j][col]]{
				return false
			}
			present[board[j][col]]= true
		}
	}

	for square := 0; square<9; square++{
		present := make(map[byte]bool)
		for i:=0;i<3;i++{
			for j:=0;j<3;j++{
				row := (square/3)*3+i
				col := (square%3)*3+j
				if board[row][col] =='.'{
					continue
				}
				if present[board[row][col]]{
					return false
				}
				present[board[row][col]]= true
			}
		}
	}
	return true
}
