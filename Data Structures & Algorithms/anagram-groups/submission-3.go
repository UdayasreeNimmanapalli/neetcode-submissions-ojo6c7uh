func groupAnagrams(strs []string) [][]string {
	var hMap = make(map[string][]string)
	for _, str := range strs{
		sortedStr := sortString(str)
		hMap[sortedStr]= append(hMap[sortedStr],str)
	}
	var res = make([][]string,0, len(hMap))
	for _, list := range hMap{
		res = append(res, list)
	}
	return res
}

func sortString(str string)string{
	strRune := []rune(str)
	sort.Slice(strRune, func(i,j int)bool{
		return strRune[i]<strRune[j]
	})
	return string(strRune)
}