func numDecodings(s string) int {
	var memo = make(map[int]int,0)
	return rec(s, 0, memo)
}

func rec(s string, index int, memo map[int]int)int{
	if index==len(s){
		return 1
	}

	if s[index]=='0'{
		return 0
	}

	if val,ok:= memo[index];ok{
		return val
	}

	res := rec(s, index+1, memo)

	if index+1<len(s){
		if s[index]=='1' || (s[index]=='2' && s[index+1]<='6'){
			res += rec(s, index+2, memo)
		}
	}

	memo[index] = res
	return res
}