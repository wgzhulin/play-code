package code

import "github.com/zhulinw/play-code/utils"

/*
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water/
*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		area := utils.MinInt(height[i], height[j]) * (j - i)
		max = utils.MaxInt(max, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
