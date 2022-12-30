// File: sfo_59i_sliding_window_maximum_s1_test.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_59i_sliding_window_maximum_s1

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {
	// ======= Test Case =======
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	// ====== Driver Code ======
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}
