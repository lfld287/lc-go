package quest

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	nnn(nums,[]int{},&res)
	return res
}

func nnn(nums []int,seq []int,res *[][]int) {
	if len(nums)==1{
		*res=append(*res,append(seq,nums[0]))
		return
	}
	Idx:=0
	reachEnd:=false
	for !reachEnd {
		curVal:=nums[Idx]
		nnn(append(nums[:Idx],nums[Idx+1:]...),append(seq,curVal),res)
		for nums[Idx]==curVal{
			if Idx==len(nums)-1{
				reachEnd=true
				break
			}
			Idx++
		}
	}
	return
}