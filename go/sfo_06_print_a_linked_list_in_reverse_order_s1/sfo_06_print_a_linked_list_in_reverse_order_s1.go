// File: sfo_06_print_a_linked_list_in_reverse_order_s1.go
// Created Time: 2022-12-31
// Author: Reanon (793584285@qq.com)

package sfo_06_print_a_linked_list_in_reverse_order_s1

import (
	. "github.com/krahets/LeetCode-Book/pkg"
)

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := recur(head)
	return res
}

func recur(node *ListNode) []int {
	// 节点为尾节点就返回
	if node.Next == nil {
		return []int{node.Val}
	}
	// 不断递归到尾节点
	res := recur(node.Next)
	// 将当前节点添加
	res = append(res, node.Val)
	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
