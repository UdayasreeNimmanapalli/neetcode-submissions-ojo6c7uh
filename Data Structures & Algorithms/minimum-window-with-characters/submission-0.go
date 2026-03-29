func minWindow(s string, t string) string {
    if len(t)>len(s){
		return ""
	}
	if t == "" || s == ""{
		return ""
	}
	var tFreq = make(map[byte]int)
	var sFreq = make(map[byte]int)
	var left = 0
	var minLen = math.MaxInt
	var minStr = ""
	for i := range t{
		tFreq[t[i]]++
	}
	var requiredChars = len(tFreq)
	var windowChars = 0
	for r := 0;r < len(s);r++{
		rightChar := s[r]
		sFreq[rightChar]++
		if count, ok := tFreq[rightChar];ok{
			if count == sFreq[rightChar]{
				windowChars++
			}
		}

		for requiredChars == windowChars && left <= r{
			strLen := len(s[left:r+1])
			if strLen < minLen{
				minLen = strLen
				minStr = s[left:r+1]
			}

			leftChar := s[left]
			if count, ok:= tFreq[leftChar];ok{
				sFreq[leftChar]--
				if sFreq[leftChar]<count{
					windowChars--
				}
			}
			left++
		}
	}
	return minStr
}
