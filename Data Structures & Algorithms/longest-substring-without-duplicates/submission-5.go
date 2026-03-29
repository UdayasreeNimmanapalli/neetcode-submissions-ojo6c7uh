func lengthOfLongestSubstring(s string) int {
	var maxLen = 0
	var l = 0
	var idxMap = make(map[byte]int)
	for r :=0; r<len(s);r++{
		if idx,ok:= idxMap[s[r]];ok{
			l = max(l, idx+1)
		}
		idxMap[s[r]]=r
		if maxLen<(r-l+1){
			maxLen = r-l+1
		}
	}
	return maxLen
}
