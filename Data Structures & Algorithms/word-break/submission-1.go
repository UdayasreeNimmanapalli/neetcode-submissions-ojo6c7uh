func wordBreak(s string, wordDict []string) bool {
    var wordMap = make(map[string]bool)
	for _,str:= range wordDict{
		wordMap[str]=true
	}
	var dp = make([]bool, len(s)+1)
	dp[0]=true
	for i:=1;i<=len(s);i++{
		for j:=0;j<i;j++{
			if wordMap[s[j:i]] && dp[j]{
				dp[i]=true
				break
			}
		}
	}
	return dp[len(s)]
}
