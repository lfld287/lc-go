package quest

import "sort"

func reorderedPowerOf2(n int) bool {

	//2的幂 那在二进制下只有一个1

	var nums []int
	for n!=0{
		tmp:=n%10
		n=n/10
		nums=append(nums,tmp)
	}
	sort.Ints(nums)
	var marker = make([]bool,len(nums))
	for idx:=range marker{
		marker[idx]=false
	}
	return backtrack869(nums,0,make([]int,len(nums)),marker)
}

func backtrack869(nums []int,idx int,seq []int,marker []bool) bool {
	var lastIdx=-1
	for i:=0;i<len(nums);i++{
		if !marker[i]{
			if lastIdx==-1||nums[lastIdx]!=nums[i]{
				seq[idx]=nums[i]
				lastIdx=i
				if idx==0&&seq[idx]==0{
					continue
				}
				marker[i]=true
				if idx==len(nums)-1{
					var tmp int
					k:=1
					for i:=len(seq)-1;i>=0;i--{
						tmp+=seq[i]*k
						k*=10
					}
					if checkIsPowerOfTwo(tmp){
						return true
					}
				}else{
					res:=backtrack869(nums,idx+1,seq,marker)
					if res{
						return true
					}
				}
				marker[i]=false
			}
		}
	}
	return false
}

func checkIsPowerOfTwo(n int)bool{
	return n&(n-1)==0
}



