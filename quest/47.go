package quest

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var marker =make([]bool,len(nums))
	for idx:=range marker{
		marker[idx]=false
	}
	nnn(nums,0,make([]int,len(nums)),&res,&marker)
	return res
}

func nnn(nums []int,idx int,seq []int,res *[][]int,marker *[]bool) {
	//search no marker
	lastIdx:=-1
	for i:=0;i<len(nums);i++{
		if !(*marker)[i]{
			if lastIdx==-1||nums[lastIdx]!=nums[i]{
				seq[idx]=nums[i]
				lastIdx=i
				(*marker)[i]=true
				if idx==len(nums)-1{
					var tmp =make([]int,len(seq))
					copy(tmp, seq)
					*res=append(*res,tmp)
				}else {
					nnn(nums,idx+1,seq,res,marker)
				}
				(*marker)[i]=false
			}
		}
	}
}