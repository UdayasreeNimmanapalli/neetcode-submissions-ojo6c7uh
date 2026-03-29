func lengthOfLongestSubstring(s string) int {
	var maxLen = 0
	for i:=0;i<len(s);i++{
		count :=0
		var charMap = make(map[byte]bool)
		for j:=i;j<len(s);j++{
			if _,ok:=charMap[s[j]];!ok{
				charMap[s[j]] = true
				count++
			}else{
				break
			}
		}
		maxLen = max(maxLen, count)
	}
	return maxLen
}
