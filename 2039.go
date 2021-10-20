package leetcode

import (
	"github.com/lfld287/lc-go/queue"
)

func networkBecomesIdle(edges [][]int, patience []int) int {
	serverNum:=len(patience)
	minPath:=make([]int,serverNum)
	for index:=range minPath{
		minPath[index]=0
	}
	//先计算每个数据服务器到主服务器的最短路径
	//长度一样，广度优先就行
	queue.New()
}