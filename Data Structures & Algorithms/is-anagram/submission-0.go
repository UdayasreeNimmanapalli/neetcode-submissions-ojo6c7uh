func isAnagram(s string, t string) bool {
	sChars := []rune(s)
	tChars := []rune(t)
	sort.Slice(sChars, func(i int, j int)bool{
		return sChars[i] < sChars[j]
	})
	sort.Slice(tChars, func(i int, j int)bool{
		return tChars[i] < tChars[j]
	})
	return string(sChars)==string(tChars)
}
