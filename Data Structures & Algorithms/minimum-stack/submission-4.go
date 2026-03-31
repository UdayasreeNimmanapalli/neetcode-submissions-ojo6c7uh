type MinStack struct {
	stack []int
 	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int,0),
		minStack: make([]int,0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minStack)==0 || (val <= this.minStack[len(this.minStack)-1]){
		this.minStack = append(this.minStack, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if (len(this.minStack)>0) && (len(this.stack)>0) && this.minStack[len(this.minStack)-1] == this.stack[len(this.stack)-1]{
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	if len(this.stack)>0{
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack)>0{
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if (len(this.minStack)>0){
		return this.minStack[len(this.minStack)-1]
	}
	return -1
}
