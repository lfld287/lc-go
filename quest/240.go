package quest

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	//收缩矩阵先
	m := len(matrix)
	n := len(matrix[0])
	var tmp [][]int
	for i := 0; i < m; i++ {
		if matrix[i][0] > target || matrix[i][n-1] < target {
			continue
		}
		tmp = append(tmp, matrix[i])
	}
	if len(tmp)==0{
		return false
	}
	m = len(tmp)
	n = len(tmp[0])

	matrix=[][]int{}
	for i := 0; i < n; i++ {
		if tmp[0][i] > target || tmp[m-1][i] < target {
			continue
		}
		var tt []int
		for j:=0;j<m;j++{
			tt=append(tt,tmp[j][i])
		}
		matrix = append(matrix, tt)
	}
	if len(matrix)==0{
		return false
	}
	m = len(matrix)
	n = len(matrix[0])

	fmt.Println(matrix)
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j]==target{
				return true
			}
		}
	}

	return false
}