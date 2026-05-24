func findOrder(numCourses int, prerequisites [][]int) []int {
	var adjList = make(map[int][]int,numCourses)
	var inDegree = make([]int,numCourses)
	var queue = make([]int,0)
	var res = make([]int,0)
	for _,edge:= range prerequisites{
		adjList[edge[1]]=append(adjList[edge[1]], edge[0])
		inDegree[edge[0]]++
	}
	for i :=0; i<numCourses; i++{
		if inDegree[i]==0{
			queue = append(queue, i)
			res = append(res, i)
		}
	}

	for len(queue)>0{
		top := queue[0]
		queue = queue[1:]
		for _, neighbor:= range adjList[top]{
			inDegree[neighbor]--
			if inDegree[neighbor]==0{
				queue = append(queue, neighbor)
				res = append(res, neighbor)
			}
		}
	}

	for i:=0;i<numCourses;i++{
		if inDegree[i]!=0{
			return []int{}
		}
	}
	return res

}
