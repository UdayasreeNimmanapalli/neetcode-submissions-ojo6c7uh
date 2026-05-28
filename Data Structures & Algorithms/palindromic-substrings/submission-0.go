func countSubstrings(s string) int {
    total := 0

	for i:=0;i<len(s);i++{
		total += rec(s, i, i)
		total += rec(s, i, i+1)
	}
	return total
}

func rec(s string, l, r int)int{
	count := 0
	for l>=0 && r<len(s) && s[l]==s[r]{
		count++
		l--
		r++
	}
	return count
}
