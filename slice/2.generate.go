package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	trigleArr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		trigleArr[i] = make([]int, i+1)
	}
	trigleArr[0][0] = 1
	if numRows == 1 {
		return trigleArr
	}
	trigleArr[1][0] = 1
	trigleArr[1][1] = 1
	for i := 2; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				trigleArr[i][j] = 1
			} else {
				trigleArr[i][j] = trigleArr[i-1][j-1] + trigleArr[i-1][j]
			}
		}
	}
	return trigleArr
}

func main() {
	fmt.Println(generate(0))
}
