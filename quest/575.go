package quest

func distributeCandies(candyType []int) int {
	tCount:=0
	tmp:=make(map[int]int)
	for _,t:=range candyType{
		if _,ok:=tmp[t];!ok{
			tCount+=1
			tmp[t]=1
		}
	}
	if tCount<len(candyType)/2{
		return tCount
	}else {
		return len(candyType)/2
	}
}
