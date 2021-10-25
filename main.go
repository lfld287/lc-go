package main

import (
	"fmt"
)

func main(){
	var money int
	fmt.Scan(&money)
	var itemCount int
	fmt.Scan(&itemCount)
	var item [][]int
	for i:=0;i<itemCount;i++{
		var val int
		var important int
		var ass int
		fmt.Scan(&val)
		fmt.Scan(&important)
		fmt.Scan(&ass)
		item = append(item,[]int{val,important,ass})
	}
	var val [][]int
	var cost [][]int
	for idx,ele:=range item{
		if ele[2]!=0{
			continue
		}
		var subItem [][]int
		for _,subEle:=range item{
			if subEle[2]== idx+1{
				subItem=append(subItem,subEle)
			}
		}
		var valList []int
		var costList []int
		valList=append(valList,ele[0]*ele[1])
		costList=append(costList,ele[0])

		if len(subItem)==0{
			valList=append(valList,ele[0]*ele[1])
			costList=append(costList,ele[0])

			valList=append(valList,ele[0]*ele[1])
			costList=append(costList,ele[0])

			valList=append(valList,ele[0]*ele[1])
			costList=append(costList,ele[0])
		}else if len(subItem)==1{
			valList=append(valList,ele[0]*ele[1]+subItem[0][0]*subItem[0][1])
			costList=append(costList,ele[0]+subItem[0][0])

			valList=append(valList,ele[0]*ele[1])
			costList=append(costList,ele[0])

			valList=append(valList,ele[0]*ele[1]+subItem[0][0]*subItem[0][1])
			costList=append(costList,ele[0]+subItem[0][0])
		}else{
			valList=append(valList,ele[0]*ele[1]+subItem[0][0]*subItem[0][1])
			costList=append(costList,ele[0]+subItem[0][0])

			valList=append(valList,ele[0]*ele[1]+subItem[1][0]*subItem[1][1])
			costList=append(costList,ele[0]+subItem[1][0])

			valList=append(valList,ele[0]*ele[1]+subItem[0][0]*subItem[0][1]+subItem[1][0]*subItem[1][1])
			costList=append(costList,ele[0]+subItem[0][0]+subItem[1][0])
		}

		val = append(val,valList)
		cost = append(cost,costList)

	}



	cnt:=len(val)

	var data =make([][]int,cnt+1)
	for idx:=range data{
		data[idx]=make([]int,money+1)
	}

	for i:=1;i<=cnt;i++{
		for c:=1;c<=money;c++{

			max := data[i-1][c]

			for j:=0;j<4;j++{
				if c-cost[i-1][j]>=0{
					tmp:=data[i-1][c-cost[i-1][j]]+val[i-1][j]
					if tmp>max{
						max=tmp
					}
				}
			}

			data[i][c]=max
		}
	}

	for idx:=range val{
		fmt.Println(val[idx],cost[idx])
	}

	for _,ele:=range data{
		fmt.Println(ele)
	}

	fmt.Println(data[cnt][money])
	return
}