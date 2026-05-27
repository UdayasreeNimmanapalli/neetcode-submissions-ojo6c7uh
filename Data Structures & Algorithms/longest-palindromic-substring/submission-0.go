func longestPalindrome(s string) string {
    if len(s)==0{
		return ""
	}
	var start int
	var end int
	for i:=0;i<len(s);i++{
		
		var maxLen = 0
		len1 := rec(s, i, i)

		len2 := rec(s, i, i+1)

		maxLen = max(maxLen, max(len1,len2))

		if maxLen>end-start+1{
			start = i-(maxLen-1)/2
			end = i + maxLen/2
		} 
	}
	return s[start:end+1]
}

func rec(s string, left, right int)int{
	for left>=0 && right<len(s)&&s[left]==s[right]{
		left--
		right++
	}
	return right-left-1
}