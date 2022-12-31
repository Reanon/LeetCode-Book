// File: sfo_06_print_a_linked_list_in_reverse_order_s2.go
// Created Time: 2022-12-31
// Author: Reanon (793584285@qq.com)

package sfo_06_print_a_linked_list_in_reverse_order_s2

import (
	"container/list"
	. "github.com/krahets/LeetCode-Book/pkg"
)

func reversePrint(head *ListNode) []int {
	// 定义时就已经初始化了
	var res []int
	stack := list.New()

	node := head
	for node != nil {
		stack.PushBack(node.Val)
		node = node.Next
	}
	for stack.Len() != 0 {
		res = append(res, stack.Back().Value.(int))
		// 移除栈末尾的值
		stack.Remove(stack.Back())
	}
	return res
}
