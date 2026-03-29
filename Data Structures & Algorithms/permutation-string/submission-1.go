func checkInclusion(s1 string, s2 string) bool {
	var s1Map = make(map[byte]int)
	var s2Map = make(map[byte]int)
	k := len(s1)
	left:=0
	for i := range s1{
		s1Map[s1[i]]++
	}
	var matches = len(s1Map)
	var foundMatches = 0
	for right :=0;right<len(s2);right++{
		rightChar := s2[right]
		s2Map[rightChar]++
		if _,ok:= s1Map[rightChar];ok{
			if s2Map[rightChar]==s1Map[rightChar]{
				foundMatches++
			}
		}
		if right-left+1>k{
			leftChar := s2[left]
			if count,ok:= s1Map[leftChar];ok{
				
				if count == s2Map[leftChar]{
					foundMatches--
				}
			}
			s2Map[leftChar]--
			left++
		}
		if foundMatches == matches{
			return true
		}
	}
	return false
}
