func isPalindrome(s string) bool {
	var l =0
	var r=len(s)-1
	for l<r{
		for l<r && !isAlphanumeric(rune(s[l])){
			l++
		}
		for l<r && !isAlphanumeric(rune(s[r])){
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])){
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphanumeric(c rune)bool{
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}