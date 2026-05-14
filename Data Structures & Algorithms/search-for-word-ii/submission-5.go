type Trie struct{
	Children [26]*Trie
	IsEnd bool
}
type Position struct{
	Row int
	Col int
}

func addWord(word string, root *Trie){
	curr:= root
	for _, char:= range word{
		index := char-'a'
		if curr.Children[index]==nil{
			curr.Children[index]= new(Trie)
		}
		curr = curr.Children[index]
	}
	curr.IsEnd = true
}

func findWords(board [][]byte, words []string) []string {
	var res []string
	var node = &Trie{}
	var visited = make(map[Position]bool)
   for _,word:= range words{
		addWord(word, node)
   }
   m:= len(board)
   n:= len(board[0])
   for row:=0;row<m;row++{
	for col:=0;col<n; col++{
		index:= board[row][col]-'a'
		if node.Children[index]!=nil{
			backtrack(row, col,m,n,"",node.Children[index],board,&res,visited)
		}
	}
   }
   return res
}

func backtrack(row int, col int, m int, n int, word string, node *Trie, board [][]byte, res *[]string, visited map[Position]bool){
	if row<0 ||col<0 || row>=m || col>=n|| visited[Position{Row:row, Col:col}]{
		return
	}
	word +=string(board[row][col])
	if node.IsEnd{
		*res = append(*res, word)
		node.IsEnd=false
	}
	
	visited[Position{row, col}]=true
	if  row+1<m && node.Children[board[row+1][col]-'a']!=nil{
		backtrack(row+1, col,m, n,word, node.Children[board[row+1][col]-'a'], board, res, visited)
	}
	if col+1<n && node.Children[board[row][col+1]-'a']!=nil {
		backtrack(row, col+1,m, n,word, node.Children[board[row][col+1]-'a'], board, res, visited)
	}
	if row-1>=0 && node.Children[board[row-1][col]-'a']!=nil {
		backtrack(row-1, col,m, n,word, node.Children[board[row-1][col]-'a'], board, res, visited)
	}
	if col-1>=0 && node.Children[board[row][col-1]-'a']!=nil{
		backtrack(row, col-1,m, n,word, node.Children[board[row][col-1]-'a'], board, res, visited)
	}
	
	visited[Position{row,col}]=false
}