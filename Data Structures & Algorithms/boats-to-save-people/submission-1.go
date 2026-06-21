func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	boats := 0
	i := 0
	j := len(people) - 1
	for i <= j {
		boats++
		if i < j && people[i]+people[j] <= limit {
			i++
		}
		j--
	}
	return boats
}