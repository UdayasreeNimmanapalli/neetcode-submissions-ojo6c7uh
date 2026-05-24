func countComponents(n int, edges [][]int) int {
    var components = 0
	var adjList = make(map[int][]int,0)
	var visited = make(map[int]bool,0)
	for _, edge := range edges{
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	for i:=0;i<n;i++{
		if !visited[i]{
			components++
			dfs(i, adjList, visited)
		}
	}
	return components
}

func dfs(i int, adjList map[int][]int, visited map[int]bool){
	visited[i] = true
	for _, neighbor := range adjList[i]{
		if !visited[neighbor]{
			dfs(neighbor, adjList, visited)
		}
	}
}