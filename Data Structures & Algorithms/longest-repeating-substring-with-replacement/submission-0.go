func characterReplacement(s string, k int) int {
	var set = make(map[byte]int)
	var maxWindow = 0
	var maxFreq = 0
	var left = 0
	for r := 0; r < len(s); r++ {
		set[s[r]]++
		maxFreq = max(maxFreq, set[s[r]])
		for (r-left+1)-maxFreq > k {
			set[s[left]]--
			left++
		}
		maxWindow = max(maxWindow, r-left+1)
	}
	return maxWindow
}