func longestPalindrome(s string) string {
	var maxLen = 0
	var longest = ""
	for i:=0;i<len(s);i++{
		for j:=i;j<len(s);j++{
			if isPalin(s, i, j){
				currLen:=j-i+1
				if currLen>maxLen{
					maxLen = currLen
					longest = s[i:j+1]
				}
			}
		}
	}
	return longest
}

func isPalin(s string, l, r int)bool{
	for l<r{
		if s[l]!=s[r]{
			return false
		}
		l++
		r--
	}
	return true
}