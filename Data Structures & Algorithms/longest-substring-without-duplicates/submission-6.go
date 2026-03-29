func lengthOfLongestSubstring(s string) int {
	var hMap = make(map[byte]bool)
	var l = 0
	var res = 0
	for r:=0;r<len(s);r++{
		for hMap[s[r]]{
			delete(hMap,s[l])
			l++
		}
		hMap[s[r]] = true
		if res<r-l+1{
			res = (r-l+1)
		}
	}
	return res
}
