func rob(nums []int) int {
   var dfs func(index int)int
   var memo = make(map[int]int,0)
   dfs = func(index int)int{
		if index>=len(nums){
			return 0
		}
		if _,ok:= memo[index];ok{
			return memo[index]
		}
		memo[index] = max(nums[index]+dfs(index+2), dfs(index+1)) 
		return memo[index]
   }
   return dfs(0)
}
