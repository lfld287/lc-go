package quest

import (
	"strings"
)

type tmp301 struct {
	l int
	r int
	start int
	val string
}

func removeInvalidParentheses(s string) []string {
	l:=0
	totalLeft:=0
	r:=0
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			l+=1
			totalLeft+=1
		}else if s[i]==')'{
			if l>0{
				l-=1
			}else {
				r+=1
			}
		}
	}
	if l==totalLeft{
		var tmp strings.Builder
		for i:=0;i<len(s);i++{
			if s[i]=='('{
				continue
			}else if s[i]==')'{
				continue
			}
			tmp.WriteRune(rune(s[i]))
		}
		return []string{tmp.String()}
	}
	var res []string
	var bfsList []tmp301
	bfsList=append(bfsList,tmp301{
		l:   l,
		r:   r,
		start: 0,
		val: s,
	})
	for len(bfsList)>0{
		ele:=bfsList[0]
		bfsList=bfsList[1:]
		if ele.l>0 {
			//not continue
			lastIdx := -1
			for i := ele.start; i < len(ele.val); i++ {
				if ele.val[i] == '(' {
					if lastIdx != -1 && lastIdx+1 == i {
						lastIdx = i
						continue
					}
					lastIdx = i
					bfsList = append(bfsList, tmp301{
						l:     ele.l - 1,
						r:     ele.r,
						start: i,
						val:   ele.val[:i] + ele.val[i+1:],
					})
				}
			}
		}
		if ele.r>0 {
			lastIdx := -1
			for i := ele.start; i < len(ele.val); i++ {
				if ele.val[i] == ')' {
					if lastIdx != -1 && lastIdx+1 == i {
						lastIdx = i
						continue
					}
					lastIdx = i
					bfsList = append(bfsList, tmp301{
						l:     ele.l,
						r:     ele.r - 1,
						start: i,
						val:   ele.val[:i] + ele.val[i+1:],
					})
				}
			}
		}
		if ele.r==0&&ele.l==0{
			if checkValid301(ele.val){
				res=append(res,ele.val)
			}
		}
	}


	return res
}

func checkValid301(s string) bool{
	l:=0
	r:=0
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			l+=1
		}else if s[i]==')'{
			if l>0{
				l-=1
			}else {
				r+=1
			}
		}
	}
	return l+r==0
}