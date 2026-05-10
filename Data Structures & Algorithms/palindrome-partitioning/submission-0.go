func partition(s string) [][]string {
	var res [][]string
	var tempStr []string
	var backtrack func(index int)
	backtrack = func(index int){
		if index==len(s){
			res = append(res, append([]string{}, tempStr...))
			return
		}
		for j:=index;j<len(s);j++{
			if isPalindrome(s[index:j+1]){
				tempStr = append(tempStr, s[index:j+1])
				backtrack(j+1)
				tempStr = tempStr[:len(tempStr)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func isPalindrome(s string)bool{
	i:=0
	j:=len(s)-1
	for i<j{
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
}