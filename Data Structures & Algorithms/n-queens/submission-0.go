func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([][]byte,n)
	for i:=0;i<n;i++{
		board[i]  = make([]byte,n)
		for j:=0;j<n;j++{
			board[i][j]='.'
		}
	}
	placeQueens(0, board, n, &res)
	return res
}

func placeQueens(col int, board [][]byte, n int, res *[][]string){
	if col == n{
		*res = append(*res, convertByteToString(board))
		return
	}

	for r :=0; r<n;r++{
		if isSafe(r, col, n, board){
			board[r][col]='Q'
			placeQueens(col+1, board, n, res)
			board[r][col]='.'
		}
	}
}

func isSafe(row int,col int,n int,board [][]byte)bool{
	// left horizontal
	tempCol := col
	for tempCol>=0{
		if board[row][tempCol]=='Q'{
			return false
		}
		tempCol--
	}

// left lower diagonal
	tempCol = col
	tempRow := row
	for tempCol>=0 && tempRow<n{
		if board[tempRow][tempCol]=='Q'{
			return false
		}
		tempRow++
		tempCol--
	}
// left upper diagonal
	tempCol = col
	tempRow = row
	for tempRow>=0 && tempCol>=0{
		if board[tempRow][tempCol]=='Q'{
			return false
		}
		tempRow--
		tempCol--
	}
	return true
}

func convertByteToString(board [][]byte)[]string{
	res := make([]string, len(board))
	for i:=0;i<len(board);i++{
		res[i]=string(board[i])
	}
	return res
}
