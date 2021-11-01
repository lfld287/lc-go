package quest

import (
	"fmt"
	"strings"
)

func removeInvalidParentheses(s string) []string {
	s=strings.TrimSpace(s)
	s=strings.Trim(s,"\n")

	offset:=0
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			offset+=1
		}else if s[i]==')'{
			offset-=1
		}

	}
}
