package quest

import "fmt"

func lengthOfLIS(nums []int) int {
	res:=make([]int,len(nums))
	for idx,val:=range nums{
		if idx==0{
			res[idx]=1
		}
		max:=0
		for j:=0;j<idx;j++{
			if val>nums[j]{
				if res[j]>max{
					max=res[j]
				}
			}
		}
		res[idx]=1+max
	}
	fmt.Println(res)
	maxValue:=0
	for _,val:=range res{
		if maxValue<val{
			maxValue=val
		}
	}

	return maxValue
}