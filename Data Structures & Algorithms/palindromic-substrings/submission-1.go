func countSubstrings(s string) int {
	var count int
	for i:=0;i<len(s);i++{
		for j:=i;j<len(s);j++{
			if isPalindrome(s, i, j){
				count++
			}
		}
	}
    return count
}

func isPalindrome(s string, left, right int)bool{
	for left<right{
		if s[left]!=s[right]{
			return false
		}
		left++
		right--
	}
	return true
}