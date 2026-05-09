func generateParenthesis(n int) []string {
	var res []string
	var str []string
	var backtrack func(open int, clos int)
	backtrack = func(open int, clos int){
		if open==n && clos==n{
			res = append(res, strings.Join(str, ""))
			return
		}
		if open<n{
			str = append(str, "(")
			backtrack(open+1, clos)
			str = str[:len(str)-1]
		}
		if clos<open{
			str = append(str,")")
			backtrack(open, clos+1)
			str = str[:len(str)-1]
		}
	}
	backtrack(0,0)
	return res
}
