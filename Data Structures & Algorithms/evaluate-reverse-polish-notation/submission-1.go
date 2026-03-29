func evalRPN(tokens []string) int {
	var stack = make([]int,0)
	for _, token := range tokens{
		switch token{
			case "+":
			a := stack[len(stack)-1]
			b:= stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := a+b
			stack = append(stack,res)
			case "-":
				a := stack[len(stack)-1]
				b:= stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				res := b-a
				stack = append(stack,res)
		case "*":
				a := stack[len(stack)-1]
				b:= stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				res := a*b
				stack = append(stack,res)
			case "/":
				a := stack[len(stack)-1]
				b:= stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				res := b/a
				stack = append(stack,res)
			default:
				num,_ := strconv.Atoi(token)
				stack = append(stack,num)
			}	
	}
	return stack[0]
}
