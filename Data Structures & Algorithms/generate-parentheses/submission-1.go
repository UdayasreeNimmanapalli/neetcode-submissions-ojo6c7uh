func generateParenthesis(n int) []string {
	var res []string
	var str []string
	var backtrack func(open int, close int)
	backtrack = func(open int, close int){
		if open == n && close == n{
			res = append(res, strings.Join(str,""))
			return
		}
		if open<n{
			str = append(str,"(")
			backtrack(open+1, close)
			str = str[:len(str)-1]
		}

		if close<open{
			str = append(str, ")")
			backtrack(open, close+1)
			str = str[:len(str)-1]
		}
	}
	backtrack(0,0)
	return res
}
