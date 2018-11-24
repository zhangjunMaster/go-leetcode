package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/**
 * 从后向前推的策略，同时注意 0和结尾
 * triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
 * 计算出每一个j的值，从中选取一个
 */
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	temp := 0
	triangle[1][0] = triangle[0][0] + triangle[1][0]
	triangle[1][1] = triangle[0][0] + triangle[1][1]
	if len(triangle) == 2 {
		return min(triangle[1][0], triangle[1][1])
	}
	for i := 2; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				triangle[i][j] = triangle[i][j] + triangle[i-1][j]
				temp = triangle[i][j]
			} else if i == j {
				triangle[i][j] = triangle[i][j] + triangle[i-1][j-1]
			} else {
				triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
			}
			if temp > triangle[i][j] {
				temp = triangle[i][j]
			}
		}
	}
	return temp
}

func main() {
	triangle := [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}
	//triangle1 := [][]int{[]int{-1}, []int{3, 2}, []int{-3, 1, -1}}

	fmt.Println(minimumTotal(triangle))
}
