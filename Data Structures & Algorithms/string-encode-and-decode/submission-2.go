type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res string
	for _, str:= range strs{
		res += strconv.Itoa(len(str))+"#"+str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	// len_str = len(encoded)
	for i := 0;i<len(encoded);{
		j:=i
		for encoded[j]!='#'{
			j++
		}
		length,_ := strconv.Atoi(string(encoded[i:j]))
		var str string
		str = encoded[j+1:length+1+j]
		res = append(res, str)
		i = length+1+j
	}
	return res
}
