func topKFrequent(nums []int, k int) []int {
	var freqMap = make(map[int]int)
	var res []int
	for _,val:= range nums{
		freqMap[val]++
	}
	var elements = make([][2]int,0)
	for num, count := range freqMap {
		elements = append(elements,[2]int{num,count})
	}
	sort.Slice(elements, func(i,j int)bool{
		return elements[i][1]>elements[j][1]
	})
	for i:=0;i<k;i++{
		res = append(res,elements[i][0])
	}
	return res
}