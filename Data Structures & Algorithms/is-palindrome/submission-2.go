func isPalindrome(s string) bool {
	var str = ""
	for _,c := range s{
		if isAlphanumeric(c){
			if unicode.IsLetter(c){
				str += strings.ToLower(string(c))
			}else{
				str += string(c)
			}
		}
	}
	l:=0
	r:=len(str)-1
	for l<r{
		if str[l]!=str[r]{
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphanumeric(char rune)bool{
	if unicode.IsLetter(char) || unicode.IsDigit(char){
		return true
	}
	return false
}