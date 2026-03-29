func isValid(s string) bool {
    var stack = make([]rune,0)
	for _, char := range s{
		switch char{
			case '[':
				stack = append(stack,']')
			case '(':
				stack = append(stack,')')
			case '{':
				stack = append(stack,'}')
			case '}',')',']':
				if len(stack)==0 || stack[len(stack)-1] != char{
					return false
				}
				stack = stack[:len(stack)-1]
		}
	}
	if len(stack)==0{
		return true
	}
	return false
}
