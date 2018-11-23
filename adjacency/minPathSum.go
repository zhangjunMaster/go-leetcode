package main

import "fmt"

/**
 * 动态规划很多都是递推，用dpn = min(dpn-1,dpn-2..) + ..
 * 有些适合从后往前，有些适合从前往后，同时满足条件后就更改状态，更改状态就是更改数组值或者另外的dp值， 其实就是dpi的值
 * 条件一般是 min max比较n n-1 n-2 ...这些的值
 * 更改状态，有的是借用外部的变量，有的是更改原先的数组
 * 大部分情况下是从后向前分析，部分情况是从前往后
 * 以为数组是 i=0,i=1,i=2时
 * 二维数组的是 i= 0，j->n时，j= 0,i->n时
 */

var arrT = [][]int{
	[]int{1, 3, 1},
	[]int{1, 5, 1},
	[]int{4, 2, 1},
}

func minPathSum(grid [][]int) int {
	n := len(grid)    //行
	m := len(grid[0]) //列
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				if i == 0 && j > 0 {
					grid[0][j] += grid[0][j-1]
				}
				if j == 0 && i > 0 {
					grid[i][0] += grid[i-1][0]
				}
			} else {
				fmt.Println(i, j)
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}
	return grid[n-1][m-1]
}

func main() {
	fmt.Println(minPathSum(arrT))
}
