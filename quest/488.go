package quest

//居然是遍历 服了

type tmp488 struct {
	seq    string
	backup []int
	insert int
}

func findMinStep(board string, hand string) int {
	var backup = make([]int, 5)

	for i := 0; i < len(hand); i++ {
		switch rune(hand[i]) {
		case 'R':
			{
				backup[0] += 1
			}
		case 'Y':
			{
				backup[1] += 1
			}
		case 'B':
			{
				backup[2] += 1
			}
		case 'G':
			{
				backup[3] += 1
			}
		case 'W':
			{
				backup[4] += 1
			}
		}
	}

	var bfsList []tmp488
	bfsList = append(bfsList, tmp488{
		seq:    board,
		backup: backup,
		insert: 0,
	})

	var removeRepeat = func(seq *string) bool {
		res:=false
		resStr:=""
		lastIdx:=0
		for i := 0; i < len(*seq);{
			repeatCount := 1
			for j := i + 1; j < len(*seq); j++ {
				if (*seq)[j] == (*seq)[i] {
					repeatCount += 1
				} else {
					break
				}
			}
			if repeatCount >= 3 {
				resStr = resStr+(*seq)[lastIdx:i]
				lastIdx=i+repeatCount
				res=true
			}
			i+=repeatCount
		}
		resStr = resStr+(*seq)[lastIdx:]
		*seq=resStr
		return res
	}

	res := -1

	for len(bfsList) > 0 {
		//fmt.Println(bfsList)
		ele := bfsList[0]
		bfsList = bfsList[1:]
		//去除重复部分
		tmpStr:=ele.seq
		for removeRepeat(&ele.seq) {
		}

		if len(ele.seq)==0{
			if res==-1||res>ele.insert{
				res=ele.insert
			}
			continue
		}else {
			ele.seq=tmpStr
		}
		repeat:=0
		for i := 0; i < len(ele.seq); i++ {
			//还得找不同，相同的差一个就行了
			if i<len(ele.seq)-1&&ele.seq[i]==ele.seq[i+1]{
				repeat+=1
				continue
			}
			if repeat>=2{
				repeat=0
				continue
			}
			//裸插是肯定没用的
			switch ele.seq[i] {
			case 'R':
				{
					if ele.backup[0] >0{
						tmpSlice:=make([]int,5)
						copy(tmpSlice,ele.backup)
						tmp:=tmp488{
							seq:    ele.seq[:i+1]+"R"+ele.seq[i+1:],
							backup: tmpSlice,
							insert: ele.insert+1,
						}
						tmp.backup[0]-=1
						bfsList=append(bfsList, tmp)
					}
				}
			case 'Y':
				{
					if ele.backup[1] >0{
						tmpSlice:=make([]int,5)
						copy(tmpSlice,ele.backup)
						tmp:=tmp488{
							seq:    ele.seq[:i+1]+"Y"+ele.seq[i+1:],
							backup: tmpSlice,
							insert: ele.insert+1,
						}
						tmp.backup[1]-=1
						bfsList=append(bfsList, tmp)
					}
				}
			case 'B':
				{
					if ele.backup[2] >0{
						tmpSlice:=make([]int,5)
						copy(tmpSlice,ele.backup)
						tmp:=tmp488{
							seq:    ele.seq[:i+1]+"B"+ele.seq[i+1:],
							backup: tmpSlice,
							insert: ele.insert+1,
						}
						tmp.backup[2]-=1
						bfsList=append(bfsList, tmp)
					}
				}
			case 'G':
				{
					if ele.backup[3] >0{
						tmpSlice:=make([]int,5)
						copy(tmpSlice,ele.backup)
						tmp:=tmp488{
							seq:    ele.seq[:i+1]+"G"+ele.seq[i+1:],
							backup: tmpSlice,
							insert: ele.insert+1,
						}
						tmp.backup[3]-=1
						bfsList=append(bfsList, tmp)
					}
				}
			case 'W':
				{
					if ele.backup[4] >0{
						tmpSlice:=make([]int,5)
						copy(tmpSlice,ele.backup)
						tmp:=tmp488{
							seq:    ele.seq[:i+1]+"W"+ele.seq[i+1:],
							backup: tmpSlice,
							insert: ele.insert+1,
						}
						tmp.backup[4]-=1
						bfsList=append(bfsList, tmp)
					}
				}
			}
			repeat=0
		}
	}

	return res
}
