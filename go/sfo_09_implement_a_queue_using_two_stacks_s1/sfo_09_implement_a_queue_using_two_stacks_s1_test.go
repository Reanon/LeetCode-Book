// File: sfo_09_implement_a_queue_using_two_stacks_s1_test.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_09_implement_a_queue_using_two_stacks_s1

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {
	// ====== Driver Code ======
	res := make([]int, 0)
	cQueue := Constructor()
	res = append(res, cQueue.DeleteHead())
	cQueue.AppendTail(5)
	cQueue.AppendTail(2)
	res = append(res, cQueue.DeleteHead())
	res = append(res, cQueue.DeleteHead())
	fmt.Println(res)
}
