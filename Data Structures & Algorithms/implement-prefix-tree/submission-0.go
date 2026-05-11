type PrefixTree struct {
	 Children [26]*PrefixTree
	 Endofword bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	curr := this
	for _, char := range word{
		index := char-'a'
		if curr.Children[index]==nil{
			curr.Children[index]=new(PrefixTree)
		}
		curr = curr.Children[index]
	}
	curr.Endofword = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this
	for _,char := range word{
		index := char-'a'
		if curr.Children[index]==nil{
			return false
		}
		curr = curr.Children[index]
	}
	return curr.Endofword
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this
	for _, char := range prefix{
		index := char-'a'
		if curr.Children[index]==nil{
			return false
		}
		curr = curr.Children[index]
	}
	return true
}
