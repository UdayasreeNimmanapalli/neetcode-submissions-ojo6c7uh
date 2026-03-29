type MinStack struct {
	Stack []int
	MStack []int
}

func Constructor() MinStack {
	return MinStack{
		Stack: make([]int,0),
		MStack: make([]int,0),
	}
}

func (this *MinStack) Push(val int) {
	if (len(this.MStack) == 0) || (this.MStack[len(this.MStack)-1]>=val){
		this.MStack = append(this.MStack, val)
	}
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	if len(this.MStack)>0 && (this.MStack[len(this.MStack)-1] == this.Stack[len(this.Stack)-1]){
		this.MStack = this.MStack[:len(this.MStack)-1]
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.Stack)>0{
		return this.Stack[len(this.Stack)-1]
	}else{
		return -1
	}
}

func (this *MinStack) GetMin() int {
	if len(this.MStack)>0{
		return this.MStack[len(this.MStack)-1]
	}else{
		return -1
	}
}
