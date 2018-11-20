package main

import "fmt"

func climbStairs(n int) int {

	number := 0
	if n == 1 {
		number = 1
		return number
	}
	if n == 2 {
		number = 2
		return number
	}
	prev := 1
	now := 2

	for i := 3; i <= n; i++ {
		next := prev + now
		prev = now
		now = next
	}
	return now
}

func main() {
	fmt.Println(climbStairs(10))
}
