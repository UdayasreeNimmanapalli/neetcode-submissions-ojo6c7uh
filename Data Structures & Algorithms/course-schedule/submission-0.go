func canFinish(numCourses int, prerequisites [][]int) bool {
    var inDegree = make([]int,numCourses)
	var queue = make([]int,0)
	var adjList = make(map[int][]int,0)
	for _, edge := range prerequisites{
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
		inDegree[edge[0]]++
	}
	for i:=0;i<numCourses;i++{
		if inDegree[i]==0{
			queue= append(queue,i)
		}
	}

	for len(queue)>0{
		top:= queue[0]
		queue=queue[1:]
		for _,n:= range adjList[top]{
			inDegree[n]--
			if inDegree[n]==0{
				queue= append(queue,n)
			}
		}
	}
	for i:=0;i<numCourses;i++{
		if inDegree[i]!=0{
			return false
		}
	}
	return true
}
