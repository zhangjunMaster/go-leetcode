package main

/*
 * 没有数组的是构建数组，形成类似于此的递推方程 fn = fn-1 + fn-2
 * 这个n是数组的下表，f是数组  s[n] = s[n-1] + s[n-2]


 */
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
