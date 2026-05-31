func dailyTemperatures(temperatures []int) []int {
	var res = make([]int,len(temperatures))
	for i:=0;i<len(temperatures);i++{
		counter:=0
		var foundWarmer = false
		for j:=i+1;j<len(temperatures);j++{
			counter++
			if temperatures[j]>temperatures[i]{
				foundWarmer = true
				break
			}
		}
		if foundWarmer{
			res[i] = counter
		}else{
			res[i] = 0
		}
		
	}
	return res
}
