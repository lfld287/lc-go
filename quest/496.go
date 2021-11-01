package quest


func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res = make([]int,len(nums1))
	for idx:=range res{
		res[idx]=-1
	}

	for idx,num:=range nums1{
		found:=false
		for j:=0;j<len(nums2);j++{
			if found&&nums2[j]>num{
				res[idx]=nums2[j]
				break
			}else if nums2[j]==num{
				found=true
			}
		}
	}

	return res
}