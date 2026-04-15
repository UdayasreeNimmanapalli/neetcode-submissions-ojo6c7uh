type LRUCache struct {
    Head, Tail *Node
	Capacity int
	HMap map[int]*Node
}

type Node struct{
	Key int
	Val int
	Prev *Node
	Next *Node
}

func newNode(key, val int)*Node{
	return &Node{
		Key:key,
		Val:val,
	}
}

func Constructor(capacity int) LRUCache {
	head := newNode(0,0)
	tail := newNode(0,0)
    head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		Head: head,
		Tail: tail,
		HMap: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
    if node,ok:= this.HMap[key];ok{
		this.deleteNode(node)
		this.addNode(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node,ok:= this.HMap[key];ok{
		this.deleteNode(node)
	}
	if len(this.HMap) == this.Capacity{
		this.deleteNode(this.Tail.Prev)
	}
	nNode := newNode(key, value)
	this.addNode(nNode)
}

func (this *LRUCache)addNode(node *Node){
	this.HMap[node.Key]= node
	oldHeadNext := this.Head.Next
	this.Head.Next = node
	node.Prev = this.Head
	node.Next = oldHeadNext
	oldHeadNext.Prev = node
}

func(this *LRUCache)deleteNode(node *Node){
	delete(this.HMap,node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}