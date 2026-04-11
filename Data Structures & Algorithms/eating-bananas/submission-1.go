func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := piles[0]
	res := math.MaxInt32
	for _,pile:= range piles{
		r = max(pile,r)
	}
	for l<=r{
		mid := (l+r)/2
		totalHours := 0
		for _,pile := range piles{
			totalHours += int(math.Ceil(float64(pile)/float64(mid)))
		}
		if totalHours <= h{
			res = min(res, mid)
			r = mid-1
		}else{
			l = mid+1
		}
	}
	return res
}
