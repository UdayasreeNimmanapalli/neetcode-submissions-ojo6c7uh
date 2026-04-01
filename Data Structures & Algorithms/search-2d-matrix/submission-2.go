func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0])==0{
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	l:= 0
	r := (rows*cols)-1

	for l<=r{
		midIndex := (l+r)/2
		midCoordinates := matrix[midIndex/cols][midIndex%cols]
		if midCoordinates ==target{
			return true
		}else if midCoordinates<target{
			l = midIndex+1
		}else{
			r = midIndex-1
		}
	}
	return false
}
