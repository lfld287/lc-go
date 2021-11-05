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

func Test47(t *testing.T){
	data:=[]int{1,1,2}
	res:=permuteUnique(data)
	fmt.Println(res)
}


func Test1218(t *testing.T){
	data:=[]int{4,12,10,0,-2,7,-8,9,-9,-12,-12,8,8}
	res:=longestSubsequence(data,0)
	fmt.Println(res)
}
