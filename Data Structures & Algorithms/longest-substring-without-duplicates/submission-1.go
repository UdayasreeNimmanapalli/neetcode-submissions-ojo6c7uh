func lengthOfLongestSubstring(s string) int {
	var maxLength = 0
	var set = make(map[byte]bool)
	var left = 0
	for r:= 0; r<len(s);r++{
		for set[s[r]]{
			delete(set, s[left])
			left++
		}
		set[s[r]]=true
		length := r-left+1
		if maxLength<length{
			maxLength = length
		}
	}
	return maxLength
}
