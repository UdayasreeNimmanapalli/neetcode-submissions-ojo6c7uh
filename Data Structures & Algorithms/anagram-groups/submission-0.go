func groupAnagrams(strs []string) [][]string {
	if len(strs)==1{
		return [][]string{strs}
	}
	
	var res = make(map[[26]int][]string)
	for _,str:= range strs{
		var freqMap[26]int
		for _,c:= range str{
			freqMap[c-'a']++
		}
		res[freqMap] = append(res[freqMap],str)
	}
	var anagrams = make([][]string,0)
	for _, group := range res{
		anagrams = append(anagrams, group)
	}
	return anagrams
}
