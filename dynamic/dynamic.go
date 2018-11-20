package main

import "fmt"

/**
 * 暴力破解，所有的数组问题都可以使用暴力破解来解决，确定begin end就可以暴力破解
 * 循环 begin
 * 循环 end
 * 从begin，end中找到 和
 */

func maxSubArray1(nums []int) int {
	length := len(nums)
	var max int = nums[0]
	for begin := 0; begin < length; begin++ {
		for end := begin; end < length; end++ {
			temp := 0
			for i := begin; i < end; i++ {
				temp += nums[i]
			}
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

/**
 * 动态规划
 * 1.不存在数组，构建数组：第一项 + 第二项 是第三项的值，从值中构建数组，获取从从第三项开始于前两项的关系，有可能是加，有可能是-
 * 2.已存在数组，计算数组：统一条件----已知数组，从0开始累加到i做一个状态,将状态与当前i比较，然后 i 累加到 j
 * 3.子过程状态 dep迭代,累加中如果小于0，则累加值dep就是当前值
 * 4.结果max 做一个变量
 * 5.如果起始值是从0开始的，则循环从1开始
 */

func maxSubArray2(nums []int) int {
	max := nums[0]
	thisSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if thisSum > 0 {
			thisSum = nums[i] + thisSum
		} else {
			thisSum = nums[i]
		}
		if max < thisSum {
			max = thisSum
		}
		fmt.Println(i, "v:", nums[i], "thisSum:", thisSum, "max:", max)
	}
	return max
}

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
