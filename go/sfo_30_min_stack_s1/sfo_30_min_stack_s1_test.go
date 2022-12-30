// File: sfo_30_min_stack_s1_test.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_30_min_stack_s1

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {
	res := make([]int, 0)
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	res = append(res, minStack.Min())
	minStack.Pop()
	res = append(res, minStack.Top())
	res = append(res, minStack.Min())
	fmt.Println(res)
}
