// File: sfo_59i_sliding_window_maximum_s1.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_59i_sliding_window_maximum_s1

func maxSlidingWindow(nums []int, k int) []int {
	var deque []int
	var res []int
	for i := 0; i < len(nums); i++ {
		// 在添加一个值到队列之前，移除前面比它小的那些值
		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, nums[i])
		// 当队内元素超过窗口大小 k 时，将队头元素移出窗口
		if i >= k && deque[0] == nums[i-k] {
			deque = deque[1:]
		}
		// 从 k-1 才开始添加结果数组
		if i >= k-1 {
			// 队头元素是队列中最大的，将队列头部的元素加入到数组中
			res = append(res, deque[0])
		}
	}
	return res
}
