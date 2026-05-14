type Trie struct{
	Children [26]*Trie
	Word string
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
	curr.Word = word
}

func findWords(board [][]byte, words []string) []string {
	var res []string
	var node = &Trie{}
   for _,word:= range words{
		addWord(word, node)
   }
   m:= len(board)
   n:= len(board[0])
   for row:=0;row<m;row++{
	for col:=0;col<n; col++{
		backtrack(row, col,m,n,node,board,&res)
	}
   }
   return res
}

func backtrack(row int, col int, m int, n int, node *Trie, board [][]byte, res *[]string){
	if row<0 ||col<0 || row>=m || col>=n{
		return
	}
	
	char := board[row][col]
	if char == '#' || node.Children[char-'a']==nil{
		return
	}

	node = node.Children[char-'a']
	if node.Word != ""{
		*res = append(*res, node.Word)
		node.Word=""
	}
	board[row][col] = '#'
	
	backtrack(row+1, col, m, n, node, board,res)
	backtrack(row-1, col, m, n, node, board,res)
	backtrack(row, col+1, m, n, node, board,res)
	backtrack(row, col-1, m, n, node, board,res)
	
	board[row][col] = char
}