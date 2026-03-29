func lengthOfLongestSubstring(s string) int {
	
	var maxLength = 0
	for i:= 0;i<len(s);i++{
		var set = make(map[byte]bool)
		length := 0
		for j:=i;j<len(s);j++{
			if set[s[j]]{
				break
			}
			set[s[j]]=true
			length++
		}
		if length>maxLength{
			maxLength = length
		}
	}
	return maxLength
}
