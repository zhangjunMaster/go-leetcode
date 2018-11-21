package main

import "fmt"

/**
 * 1.从左到右的一次循环
 * 2.寻找循环到i就一个子过程结束，从新获取值，状态值从新获取
 * 3.dp(0),dp(1),dp(2),dp(3),dp(4) ...dp(n) 与i的关系，min,max，i可以变化
 * 递推是连续的，递推是当前的num[i]与以前的dp关系和以前的dp的关系。
 * 例如：[5,4,8,1]
 * dp(0) = 5 初始值
 * dp(1) = 5 初始值
 * dp(2) = max(num[2]+dp(0),dp(1))
 * dp(3) = max(num[3]+dp(1),dp(2))
 */

func rob(nums []int) int {
	var dp int
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		dp = nums[0]
	}
	if len(nums) > 1 {
		if nums[1] > nums[0] {
			dp = nums[1]
		} else {
			dp = nums[0]
		}
	}
	// 初始化计算的值
	dp_2 := nums[0]
	dp_1 := dp
	for i := 2; i < len(nums); i++ {
		// dp状态更新
		dp = nums[i] + dp_2
		if nums[i]+dp_2 > dp_1 {
			dp = nums[i] + dp_2
		} else {
			dp = dp_1
		}
		// 递推值变化
		dp_2 = dp_1
		dp_1 = dp
	}
	return dp
}

func main() {
	max := rob([]int{2, 7, 9, 3, 1})
	fmt.Println(max)
}
