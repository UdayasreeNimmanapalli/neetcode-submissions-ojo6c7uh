func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1)+len(nums2)
	A := nums1
	B := nums2
	if len(B)<len(A){
		A,B = B,A
	}
	l := 0
	r := len(A)
	for l<=r{
		leftPart := (l+r)/2
		rightPart := (totalLength+1)/2 - (leftPart)

		Aleft := math.MinInt64
		if leftPart>0{
			Aleft = A[leftPart-1]
		}
		Aright := math.MaxInt64
		if leftPart < len(A){
			Aright = A[leftPart]
		}

		Bleft := math.MinInt64
		if rightPart>0{
			Bleft = B[rightPart-1]
		}
		Bright := math.MaxInt64
		if rightPart<len(B){
			Bright = B[rightPart]
		}
		if Aleft<=Bright && Bleft<=Aright{
			if totalLength % 2 != 0{
				return float64(max(Aleft,Bleft))
			}
			return (float64(max(Aleft,Bleft))+float64(min(Aright,Bright)))/2.0
		}else if Aleft > Bright{
			r = leftPart-1
		}else{
			l = leftPart+1
		}
	}
	return -1
}
