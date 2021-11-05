package quest

import "fmt"

func longestSubsequence(arr []int, difference int) int {
	m:=make(map[int]int)
	res:=0
	for i:=0;i<len(arr);i++{
		max:=1
		if v,ok:=m[arr[i]-difference];ok{
			if v+1>max{
				max=v+1
			}
		}

		if _,ok:=m[arr[i]];ok{
			if max>m[arr[i]]{
				m[arr[i]]=max
			}
		}else{
			m[arr[i]]=max
		}

		if m[arr[i]]>res{
			res = m[arr[i]]
		}
	}
	fmt.Println(m)
	return res
}