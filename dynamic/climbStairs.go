package main

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
	// 子过程状态更改，使用next做临时变量
	for i := 3; i <= n; i++ {
		next := prev + now
		prev = now
		now = next
	}
	return now
}
