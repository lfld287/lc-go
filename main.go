package main

import (
	"fmt"
	"math"
	"sort"
)

func testStackPointer(p *int) {
	k := 3
	p = &k
}
type mData struct{
	stringData []string
	valData []int64
}

func (p *mData)Len() int{
	return len(p.stringData)
}

func (p *mData)Less(i,j int)bool{
	return p.valData[i]<p.valData[j]
}

func (p *mData)Swap(i,j int){
	tmp:=p.valData[i]
	p.valData[i]=p.valData[j]
	p.valData[j]=tmp
	tmpS:=p.stringData[i]
	p.stringData[i]=p.stringData[j]
	p.stringData[j]=tmpS
}

func main() {
	var data mData
	data=mData{
		stringData: make([]string,20),
		valData:    make([]int64,20),
	}
	var input int64
	fmt.Scan(&input)
	var res []int64
	maxNum:=int64(math.Ceil(math.Sqrt(float64(input))))
	var i int64
	for i=2;i<=maxNum;i++{
		for  input%i==0{
			res = append(res,i)
			input=input/i
		}
	}
	if input>1{
		res = append(res,input)
	}
	for _,ele:=range res{
		fmt.Printf("%d ",ele)
	}

	fmt.Println()
}