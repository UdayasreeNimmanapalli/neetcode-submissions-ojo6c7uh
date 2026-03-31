func evalRPN(tokens []string) int {
	var stack = make([]int,0)
	for _, token:= range tokens{
		switch token{
			case "+":
				b:= stack[len(stack)-1]
				a:= stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, a+b)
			case "-":
				b:= stack[len(stack)-1]
				a:= stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, a-b)
			case "/":
				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, a/b)
			case "*":
				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, a*b)
			default:
				num,_:= strconv.Atoi(token)
				stack = append(stack,num)
		}
	}
	return stack[0]
}
