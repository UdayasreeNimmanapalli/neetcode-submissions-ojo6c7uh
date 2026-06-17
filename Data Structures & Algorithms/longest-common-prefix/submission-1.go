func longestCommonPrefix(strs []string) string {
    var res = ""
	res = strs[0]
	for i:=1;i<len(strs);i++{
		j:=0
		minLen:=len(res)
		if len(strs[i])<minLen{
			minLen = len(strs[i])
		}
		for j<minLen&&res[j]==strs[i][j]{
			j++
		}
		res = res[:j]
		if res == ""{
			break
		}
	}
	return res
}
