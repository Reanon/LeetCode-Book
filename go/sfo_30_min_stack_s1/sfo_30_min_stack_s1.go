// File: sfo_30_min_stack_s1.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_30_min_stack_s1

import "container/list"

type MinStack struct {
	stack *list.List
	// 辅助栈
	minStack *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		// 初始化 List
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	// 添加元素
	this.stack.PushBack(x)
	// 向最小栈中添加元素
	if this.minStack.Len() == 0 || this.minStack.Back().Value.(int) >= x {
		this.minStack.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	e := this.stack.Back()
	m := this.minStack.Back()
	// 如果最小栈的元素等于当前弹出的元素值，则最小栈也要弹出
	if e.Value.(int) == m.Value.(int) {
		this.minStack.Remove(m)
	}
	this.stack.Remove(e)
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.minStack.Back().Value.(int)
}
