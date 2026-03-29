func topKFrequent(nums []int, k int) []int {
	var freqMap = make(map[int]int)
	var res []int
	for _,val:= range nums{
		freqMap[val]++
	}
	var elements []Elements
	for num, count := range freqMap {
		elements = append(elements, Elements{Key:num,Value:count})
	}
	sort.Slice(elements, func(i,j int)bool{
		return elements[i].Value>elements[j].Value
	})
	for i:=0;i<k;i++{
		res = append(res,elements[i].Key)
	}
	return res
}
type Elements struct{
	Key int
	Value int
}