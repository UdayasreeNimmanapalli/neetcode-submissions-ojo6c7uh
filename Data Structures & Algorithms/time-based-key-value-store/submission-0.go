type TimeMap struct {
	timeMap map[string]map[int][]string
}

func Constructor() TimeMap {
	return TimeMap{
		timeMap : make(map[string]map[int][]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _,ok:=this.timeMap[key];!ok{
		this.timeMap[key]= make(map[int][]string)
	}
	this.timeMap[key][timestamp]=append(this.timeMap[key][timestamp],value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _,ok:= this.timeMap[key];!ok{
		return ""
	}
	maxTime := 0
	for time:= range this.timeMap[key]{
		if time <= timestamp{
			maxTime = max(maxTime, time)
		}
	}
	if maxTime == 0{
		return ""
	}
	values := this.timeMap[key][maxTime]
	return values[len(values)-1]
}
