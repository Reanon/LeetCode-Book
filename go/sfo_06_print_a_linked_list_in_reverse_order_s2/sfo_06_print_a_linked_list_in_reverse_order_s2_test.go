// File: sfo_06_print_a_linked_list_in_reverse_order_s2_test.go
// Created Time: 2022-12-31
// Author: Reanon (793584285@qq.com)

package sfo_06_print_a_linked_list_in_reverse_order_s2

import (
	"fmt"
	"testing"

	. "github.com/krahets/LeetCode-Book/pkg"
)

func TestReversePrint(t *testing.T) {
	// ======= Test Case =======
	head := ArrayToLinkedList([]int{1, 3, 2})
	// ====== Driver Code ======
	res := reversePrint(head)
	fmt.Println(res)
}
