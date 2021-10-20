package leetcode

import "fmt"

func networkBecomesIdle(edges [][]int, patience []int) int {
	serverNum := len(patience)

	//先计算每个数据服务器到主服务器的最短路径
	//长度一样，广度优先就行
	//不能广度优先 用那个

	confirmed := make([]bool, serverNum)
	for index := range confirmed {
		confirmed[index] = false
	}

	dis := make([]int, serverNum)
	for index := range dis {
		dis[index] = -1
	}
	dis[0] = 0

	for {
		//all confirmed?
		allConfirmed := true
		for _, ele := range confirmed {
			if ele == false {
				allConfirmed = false
				break
			}
		}
		if allConfirmed {
			break
		}

		currentIndex := -1
		for index := range dis {

			if confirmed[index] == true ||dis[index]==-1{
				continue
			}

			if currentIndex == -1{
				currentIndex = index
			} else if dis[currentIndex]==-1{
				currentIndex = index
			} else if dis[currentIndex] > dis[index] {
				currentIndex = index
			}
		}

		if currentIndex == -1 {
			break
		}

		confirmed[currentIndex] = true

		for _, ele := range edges {
			if ele[0] == currentIndex {
				if dis[ele[1]] > dis[currentIndex]+1 || dis[ele[1]] == -1 {
					dis[ele[1]] = dis[currentIndex] + 1
				}
			}else if ele[1]==currentIndex{
				if dis[ele[0]] > dis[currentIndex]+1 || dis[ele[0]] == -1 {
					dis[ele[0]] = dis[currentIndex] + 1
				}
			}
		}
	}

	fmt.Println(dis)

	maxCost := 0

	for index := range dis {
		if patience[index] == 0 {
			continue
		}
		sendTime := dis[index] * 2

		count := sendTime / patience[index]
		if count > 1 {
			sendTime = patience[index]*(count-1) + sendTime
		}

		if maxCost < sendTime {
			maxCost = sendTime
		}

	}

	return maxCost + 1
}
