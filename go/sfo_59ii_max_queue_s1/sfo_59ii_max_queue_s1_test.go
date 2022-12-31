// File: sfo_59ii_max_queue_s1_test.go
// Created Time: 2022-12-31
// Author: Reanon (793584285@qq.com)

package sfo_59ii_max_queue_s1

import (
	"fmt"
	"testing"
)

func TestMaxQueue(t *testing.T) {
	// ====== Driver Code ======
	var res []int
	maxQueue := Constructor()
	maxQueue.Push_back(1)
	maxQueue.Push_back(2)
	res = append(res, maxQueue.Max_value())
	res = append(res, maxQueue.Pop_front())
	res = append(res, maxQueue.Max_value())
	fmt.Println(res)
}
