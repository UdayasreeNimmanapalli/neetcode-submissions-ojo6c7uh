type Twitter struct {
    tweetMap map[int][]tweet
	followerMap map[int]map[int]bool
	time int
}

type tweet struct{
	tweetID int
	time int
}

func Constructor() Twitter {
    return Twitter{
		tweetMap: make(map[int][]tweet,0),
		followerMap: make(map[int]map[int]bool,0),
		time: 0,
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.tweetMap[userId] = append(this.tweetMap[userId], tweet{tweetID: tweetId, time: this.time})
	this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    var mh = &mHeap{}
	heap.Init(mh)
	top10Tweet := func(uid int){
		tweets := this.tweetMap[uid]
		n:= len(tweets)
		for i:=n-1;i>=0 && i>=n-10;i--{
			heap.Push(mh, tweets[i])
		}
	}

	top10Tweet(userId)

	if followees,ok:=this.followerMap[userId];ok{
		for followee:= range followees{
			top10Tweet(followee)
		}
	}
	result := make([]int,0)
	for mh.Len()>0 && len(result)<10{
		top:=heap.Pop(mh).(tweet)
		result = append(result, top.tweetID)
	}
	return result
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if _,ok:=this.followerMap[followerId];!ok{
		this.followerMap[followerId]=make(map[int]bool,0)
	}
	this.followerMap[followerId][followeeId]=true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.followerMap[followerId], followeeId)
}

type mHeap []tweet

func(m mHeap)Len()int{
	return len(m)
}

func(m mHeap)Less(i, j int)bool{
	return m[i].time>m[j].time
}

func(m mHeap)Swap(i, j int){
	m[i],m[j]= m[j],m[i]
}

func(m *mHeap)Push(x interface{}){
	*m = append(*m, x.(tweet))
}

func(m *mHeap)Pop()interface{}{
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}
