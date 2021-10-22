package quest

func networkBecomesIdle(edges [][]int, patience []int) int {
	serverNum := len(patience)
	//先计算每个数据服务器到主服务器的最短路径
	//长度一样，广度优先就行
	//不能广度优先 用那个

	var tmp = make([][]int, serverNum)
	for _, v := range edges {
		d1 := v[0]
		d2 := v[1]
		tmp[d1] = append(tmp[d1], d2)
		tmp[d2] = append(tmp[d2], d1)
	}
	edges = tmp

	dis := make([]int, serverNum)
	for index := range dis {
		dis[index] = -1
	}
	dis[0] = 0

	//dij 太慢了 似乎，还得bfs 麻烦

	var q []int
	q = append(q, 0)
	for len(q) > 0 {
		ele := q[0]
		q = q[1:]
		for _, edge := range edges[ele] {
			if dis[edge] != -1 {
				continue
			}
			dis[edge] = dis[ele] + 1
			q = append(q, edge)
		}
	}

	maxCost := 0

	for index := range dis {
		if patience[index] == 0 {
			continue
		}
		sendTime := dis[index] * 2

		//这里还是搞不懂
		//减一是因为到达的那一秒
		count := (sendTime - 1) / patience[index]
		sendTime = patience[index]*(count) + sendTime

		if maxCost < sendTime {
			maxCost = sendTime
		}

	}

	return maxCost + 1
}
