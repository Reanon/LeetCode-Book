// File: sfo_59ii_max_queue_s1.go
// Created Time: 2022-12-31
// Author: Reanon (793584285@qq.com)

package sfo_59ii_max_queue_s1

type MaxQueue struct {
	// 使用切片模拟队列
	queue []int
	deque []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	// 加入双端队列前，需要把小于 value 的元素移出
	for len(this.deque) > 0 && this.deque[len(this.deque)-1] < value {
		this.deque = this.deque[:len(this.deque)-1]
	}
	this.deque = append(this.deque, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	val := this.queue[0]
	// 移出元素前，需要判断当前最大值是否为该元素
	if val == this.deque[0] {
		this.deque = this.deque[1:]
	}
	this.queue = this.queue[1:]
	return val
}
