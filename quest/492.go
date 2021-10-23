package quest

import (
	"bufio"
	"fmt"
	"os"
)


func constructRectangle(area int) []int {
	left:=0
	right:=area
	middle:=1
	for right-left>1{
		middle=(left+right)/2
		if middle*middle>area{
			right = middle
		}else if middle*middle<area{
			left = middle
		}else {
			break
		}
	}
	root:=middle
	l:=root
	w:=root
	for{
		tmpArea:=l*w
		if tmpArea<area{
			l+=1
		}else if tmpArea>area{
			w-=1
		}else {
			break
		}
	}

	fmt.Println(root,l,w,area)
	return []int{l,w}
}