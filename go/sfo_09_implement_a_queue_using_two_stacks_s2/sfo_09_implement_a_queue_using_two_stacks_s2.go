// File: sfo_09_implement_a_queue_using_two_stacks_s1.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_09_implement_a_queue_using_two_stacks_s2

type CQueue struct {
	// inStack 用于存储入队元素
	inStack []int
	// outStack 用于存储出队元素
	outStack []int
}

func Constructor() CQueue {
	// 初始化
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	// 添加队尾: 入队
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	// 如果 outStack 数组为空
	if len(this.outStack) == 0 {
		// 如果入栈不为空就将 inStack 中所有元素放到 outStack
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	// 如果出栈中有值，则返回栈顶元素
	if len(this.outStack) > 0 {
		val := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		// 返回 outStack 栈顶元素
		return val
	}
	// 队列中没有元素，返回 -1
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
