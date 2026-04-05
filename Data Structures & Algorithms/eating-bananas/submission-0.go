func minEatingSpeed(piles []int, h int) int {
	var minSpeed = 0
	l:=1
	r:=piles[0]
	for _,pile := range piles{
		r = max(r,pile)
	}
	for l<=r{
		mid := (l+r)/2
		var time = 0
		for _,pile := range piles{
			time += int(math.Ceil(float64(pile)/float64(mid)))
		}
		if time<=h{
			minSpeed = mid
			r = mid-1
		}else{
			l = mid+1
		}
	}
	return minSpeed
}
