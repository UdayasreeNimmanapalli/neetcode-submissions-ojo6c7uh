func isAnagram(s string, t string) bool {
	var hMap = make(map[[26]int]string)
	charCountS := [26]int{}
	charCountT := [26]int{}

	for _, c:= range s{
		charCountS[c-'a']++
	}
	hMap[charCountS] = s

	for _, c:= range t{
		charCountT[c-'a']++
	}
	hMap[charCountT] = t
	return hMap[charCountS] == hMap[charCountT]
}
