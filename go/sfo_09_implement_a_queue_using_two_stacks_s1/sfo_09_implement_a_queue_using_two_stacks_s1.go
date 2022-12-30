// File: sfo_09_implement_a_queue_using_two_stacks_s1.go
// Created Time: 2022-12-30
// Author: Reanon (793584285@qq.com)

package sfo_09_implement_a_queue_using_two_stacks_s1

import (
	// 导入内置的 list 包
	"container/list"
)

type CQueue struct {
	// A 用于存储正序元素
	A *list.List
	// B 用于存储逆序元素
	B *list.List
}

func Constructor() CQueue {
	// 初始化
	return CQueue{
		A: list.New(),
		B: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	// 添加队尾: 入队
	this.A.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 对 list 的长度进行判断
	if this.B.Len() == 0 {
		for this.A.Len() > 0 {
			// 将 A 中的值逆序放入 B 中
			this.B.PushBack(this.A.Remove(this.A.Back()))
		}
	}
	if this.B.Len() > 0 {
		e := this.B.Back()
		this.B.Remove(e)
		// 返回元素中的值
		return e.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
