type LRUCache struct {
    Head, Tail *Node
	HMap map[int]*Node
	Capacity int
}

type Node struct{
	Prev *Node
	Next *Node
	Val int
	Key int
}

func newNode(key, val int)*Node{
	return &Node{
		Key: key,
		Val: val,
	}
}
func lruCache(head, tail *Node, capacity int)LRUCache{
	return LRUCache{
		Head: head,
		Tail: tail,
		HMap: make(map[int]*Node),
		Capacity: capacity,
	}
}

func Constructor(capacity int) LRUCache {
   head:= newNode(0,0)
   tail := newNode(0,0)
   head.Next = tail
   tail.Prev = head

   //return lruCache(head, tail, capacity)

   return  LRUCache{
		Head: head,
		Tail: tail,
		HMap: make(map[int]*Node),
		Capacity: capacity,
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
    if node, ok:=this.HMap[key];ok{
		this.deleteNode(node)
	}
	if len(this.HMap)==this.Capacity{
		this.deleteNode(this.Tail.Prev)
	}
	this.addNode(newNode(key, value))
}

func(this *LRUCache)addNode(node *Node){
	this.HMap[node.Key]=node
	oldNext := this.Head.Next
	this.Head.Next = node
	node.Next = oldNext
	node.Prev = this.Head
	oldNext.Prev = node
}

func(this *LRUCache)deleteNode(node *Node){
	delete(this.HMap, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev =node.Prev
}
