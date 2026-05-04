func lengthOfLongestSubstring(s string) int {
	var hMap = make(map[byte]bool)
	l:=0
	var maxCount = 0
	for r:=0;r<len(s);r++{
		if _,ok:=hMap[s[r]];!ok{
			hMap[s[r]]=true
			maxCount = max(maxCount,r-l+1)
		}else{
			for hMap[s[r]]{
				delete(hMap, s[l])
				l++
			}
			hMap[s[r]]=true
			maxCount = max(maxCount,r-l+1)
		}
	}
	return maxCount
}
