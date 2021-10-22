package quest

import (
	"fmt"
	"sort"
)

// 摩尔投票法
func majorityElement(nums []int) []int {
	cnt := len(nums)
	limit := cnt / 3 //最多只可能有两个数，第三个数开始就肯定小于了
	vote1 := 0
	vote2 := 0
	e1 := 0
	e2 := 0
	for _, num := range nums {
		if vote1 > 0 && num == e1 {
			vote1 += 1
		} else if vote2 > 0 && num == e2 {
			vote2 += 1
		} else if vote1 == 0 {
			vote1 += 1
			e1 = num
		} else if vote2 == 0 {
			vote2++
			e2 = num
		} else {
			vote1 -= 1
			vote2 -= 1
		}
	}

	var res []int

	if vote1 > 0 {
		count:=0
		for _,num:=range nums{
			if e1==num{
				count+=1
			}
		}
		if count>limit{
			res = append(res, e1)
		}
	}

	if vote2 > 0 {
		count:=0
		for _,num:=range nums{
			if e2==num{
				count+=1
			}
		}
		if count>limit{
			res = append(res, e2)
		}
	}

	fmt.Println(e1, vote1, e2, vote2)

	return res
}

func majorityElementNative(nums []int) []int {
	cnt := len(nums)
	limit := cnt / 3 //最多只可能有两个数，第三个数开始就肯定小于了
	var res []int
	sort.Ints(nums)
	cur := 0
	for idx := 0; idx <= cnt-1; idx++ {

		if nums[cur] != nums[idx] {
			if idx-cur > limit {

				res = append(res, nums[cur])
			}
			cur = idx
		}
		if idx == cnt-1 {
			if idx-cur+1 > limit {
				res = append(res, nums[cur])
			}
		}
	}

	return res
}
