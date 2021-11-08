package quest

import "strconv"

func getHint(secret string, guess string) string {
	l:=len(secret)
	cows:=0
	bulls:=0
	m:=make(map[uint8]int)
	for i:=0;i<l;i++{
		if _,ok:=m[guess[i]];ok{
			m[guess[i]]+=1
		}else{
			m[guess[i]]=1
		}
	}
	for i:=0;i<l;i++{
		if secret[i]==guess[i]{
			bulls+=1
			m[secret[i]]-=1
		}
	}
	for i:=0;i<l;i++{
		if secret[i]==guess[i]{
			continue
		}else{
			if _,ok:=m[secret[i]];ok{
				if m[secret[i]]>0{
					cows+=1
					m[secret[i]]-=1
				}
			}
		}
	}
	resStr:=strconv.Itoa(bulls)+"A"+strconv.Itoa(cows)+"B"
	return resStr
}