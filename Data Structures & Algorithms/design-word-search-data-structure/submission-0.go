type WordDictionary struct {
    Children map[rune]*WordDictionary
	IsEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{
		Children: make(map[rune]*WordDictionary),
		IsEnd: false,
	}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this
	for _, char := range word{
		if curr.Children[char] ==nil{
			curr.Children[char] = &WordDictionary{
                Children: make(map[rune]*WordDictionary),
                IsEnd: false,
            }
		}
		curr = curr.Children[char]
	}
	curr.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(index int, node *WordDictionary)bool
	dfs = func(index int, node *WordDictionary)bool{
		if index == len(word){
			return node.IsEnd
		}
		char := rune(word[index])

		if char=='.'{
			for _, child := range node.Children{
				if dfs(index+1, child){
					return true
				}
			}
			return false
		}else{
			if node.Children[char]==nil{
				return false
			}
			return dfs(index+1, node.Children[char])
		}
	}
	return dfs(0, this)
}
