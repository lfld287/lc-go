package quest

func isValid(s string) bool {
	l1:=0
	r1:=0
	l2:=0
	r2:=0
	l3:=0
	r3:=0
	var lastLeft []int
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			l1+=1
			lastLeft=append(lastLeft,1)
		}else if s[i]==')'{
			if len(lastLeft)==0||lastLeft[len(lastLeft)-1]!=1{
				return false
			}else {
				lastLeft=lastLeft[:len(lastLeft)-1]
			}
			if l1>0{
				l1-=1
			}else{
				return false
			}
		}else if s[i]=='['{
			l2+=1
			lastLeft=append(lastLeft,2)
		}else if s[i]==']'{
			if len(lastLeft)==0||lastLeft[len(lastLeft)-1]!=2{
				return false
			}else {
				lastLeft=lastLeft[:len(lastLeft)-1]
			}
			if l2>0{
				l2-=1
			}else{
				return false
			}
		}else if s[i]=='{'{
			l3+=1
			lastLeft=append(lastLeft,3)
		}else if s[i]=='}'{
			if len(lastLeft)==0||lastLeft[len(lastLeft)-1]!=3{
				return false
			}else {
				lastLeft=lastLeft[:len(lastLeft)-1]
			}
			if l3>0{
				l3-=1
			}else{
				return false
			}
		}
	}

	return l1+r1+l2+r2+l3+r3==0
}
