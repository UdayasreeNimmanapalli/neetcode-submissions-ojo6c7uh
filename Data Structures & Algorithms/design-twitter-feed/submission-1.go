type Twitter struct {
    tweet []Tweet
	followerID map[int]map[int]bool
	timestamp int
}

type Tweet struct{
	userID int
	tweetID int
	timestamp int
}
func Constructor() Twitter {
    return Twitter{
		tweet : make([]Tweet,0),
		followerID: make(map[int]map[int]bool),
		timestamp: 0,
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.timestamp++
    this.tweet = append(this.tweet, Tweet{userID: userId, tweetID:tweetId, timestamp:this.timestamp})
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    var hp = &mHeap{}
	heap.Init(hp)
	for _, tweet := range this.tweet{
		if tweet.userID == userId || this.followerID[userId][tweet.userID]{
			if hp.Len()<10{
				heap.Push(hp, tweet)
			}else{
				if tweet.timestamp > (*hp)[0].timestamp{
					heap.Pop(hp)
					heap.Push(hp, tweet)
				}
			}								
		}
	}
	var res []int
	for hp.Len()>0{
		res = append(res, heap.Pop(hp).(Tweet).tweetID)
	}

	// Reverse to get descending order
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
	return res
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if this.followerID[followerId]==nil{
		this.followerID[followerId] = make(map[int]bool,0)
	}
	this.followerID[followerId][followeeId]=true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.followerID[followerId],followeeId)
}

type mHeap []Tweet

func(m mHeap)Len()int{
	return len(m)
}

func(m mHeap)Less(i, j int)bool{
	return m[i].timestamp<m[j].timestamp
}

func(m mHeap)Swap(i, j int){
	m[i],m[j]= m[j],m[i]
}

func(m *mHeap)Push(x interface{}){
	*m = append(*m, x.(Tweet))
}

func(m *mHeap)Pop()interface{}{
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}