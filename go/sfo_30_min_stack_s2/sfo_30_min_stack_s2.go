// File: sfo_30_min_stack_s1.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_30_min_stack_s2

type MinStack struct {
	stack []int
	// 辅助栈
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	// 添加元素
	this.stack = append(this.stack, x)
	// 向最小栈中添加元素
	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= x {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	e := this.stack[len(this.stack)-1]
	m := this.minStack[len(this.minStack)-1]
	// 如果最小栈的元素等于当前弹出的元素值，则最小栈也要弹出
	if e == m {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}
