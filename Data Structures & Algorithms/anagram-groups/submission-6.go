func groupAnagrams(strs []string) [][]string {
	var res = make([][]string,0)
	var hMap = make(map[[26]int][]string)
	for _,str:= range strs{
		charCount := [26]int{}
		for _,char := range str{
			charCount[char-'a']++
		}
		hMap[charCount]=append(hMap[charCount], str)
	}
	for _, str := range hMap{
		res = append(res, str)
	}
	return res
}
