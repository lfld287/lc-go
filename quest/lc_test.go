package quest

import (
	"fmt"
	"testing"
)

func Test492(t *testing.T){
	constructRectangle(9999998)
}

func Test240(t *testing.T){
	var data [][]int
	data=[][]int{
	{-5},
	}
	searchMatrix(data,
	20)
}

func Test301(t *testing.T){
	s:="())(((()m)("
	res:=removeInvalidParentheses(s)
	fmt.Println(res)
}