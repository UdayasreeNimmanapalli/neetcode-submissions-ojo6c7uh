func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 2 pointer method
	len1 := len(nums1)
	len2 := len(nums2)
	l:=0
	r:=0
	counter := (len(nums1)+len(nums2))/2
	var prev , curr int
	for i:=0;i<=counter;i++{
		prev=curr
		if l<len1 && r<len2{
			if nums1[l]<nums2[r]{
				curr=nums1[l]
				l++
			}else{
				curr = nums2[r]
				r++
			}
		}else if l<len1{
			curr=nums1[l]
			l++
		}else{
			curr = nums2[r]
			r++
		}
	}
	if (len1+len2)%2==0{
		return float64(prev+curr)/2.0
	}
	return float64(curr)

}
