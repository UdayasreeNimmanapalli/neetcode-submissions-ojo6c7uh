func topKFrequent(nums []int, k int) []int {
	var res = make([]int,0)
	var fMap = make(map[int]int)
	var slice = make([][2]int,0)
	for _,num := range nums{
		fMap[num]++
	}
	for num, count := range fMap{
		slice = append(slice, [2]int{num, count})
	}
	sort.Slice(slice, func(i, j int)bool{
		return slice[i][1]>slice[j][1]
	})
	for i:= 0; i<k ; i++{
		res = append(res,slice[i][0])
	}
	return res
}