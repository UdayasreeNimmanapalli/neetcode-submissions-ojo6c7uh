func lengthOfLongestSubstring(s string) int {
	var hMap = make(map[byte]int)
	var res = 0
	l:=0
	for r:=0;r<len(s);r++{
		if idx,ok:=hMap[s[r]];ok{
			l = max(l, idx+1)
		}
		hMap[s[r]]=r
		if res<r-l+1{
			res = r-l+1
		}
	}
	return res
}
