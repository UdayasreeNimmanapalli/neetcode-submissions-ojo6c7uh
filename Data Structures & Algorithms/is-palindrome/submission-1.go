func isPalindrome(s string) bool {
	var str = ""
	for _, char := range s{
		if isAlphnumeric(char){
			if unicode.IsLetter(char){
			str += strings.ToLower(string(char))
			}else{
				str += string(char)
			}
		}
	}

	l:=0
	r:= len(str)-1
	for l<r{
		if str[l]!=str[r]{
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphnumeric(c rune)bool{
	if unicode.IsLetter(c) || unicode.IsDigit(c){
		return true
	}
	return false
}