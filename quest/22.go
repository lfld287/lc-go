package quest

import "strings"

func generateParenthesis(n int) []string {
	var res []string
	backtrack22(0,n,make([]rune,2*n),0,2*n,&res)
	return res
}


func backtrack22(availableRight int,availableLeft int,seq []rune,idx int,total int,res *[]string){
	if availableLeft>0{
		seq[idx]='('
		if idx==total-1{
			var b strings.Builder
			for i:=0;i<len(seq);i++{
				b.WriteRune(seq[i])
			}
			*res=append(*res,b.String())
		}else {
			backtrack22(availableRight+1,availableLeft-1,seq,idx+1,total,res)
		}
	}
	if availableRight>0{
		seq[idx]=')'
		if idx==total-1{
			var b strings.Builder
			for i:=0;i<len(seq);i++{
				b.WriteRune(seq[i])
			}
			*res=append(*res,b.String())
		}else {
			backtrack22(availableRight-1,availableLeft,seq,idx+1,total,res)
		}
	}
}