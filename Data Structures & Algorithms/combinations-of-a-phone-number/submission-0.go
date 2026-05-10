func letterCombinations(digits string) []string {
	if len(digits)==0{
		return []string{}
	}
	var res []string
	digitMap := map[string]string{
		"2":"abc",
		"3":"def",
		"4":"ghi",
		"5":"jkl",
		"6":"mno",
		"7":"pqrs",
		"8":"tuv",
		"9":"wxyz",
	}
	var backtrack func(index int, currStr string)
	backtrack = func(index int, currStr string){
		if index == len(digits){
			res = append(res, currStr)
			return
		}
		digit := string(digits[index])
		chars := digitMap[digit]
		for _, char:= range chars{
			backtrack(index+1, currStr+string(char))
		}
	}
	backtrack(0, "")
	return res
}
