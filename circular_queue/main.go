package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	k := 5
	obj := Constructor(k)
	param_1 := obj.EnQueue(1)
	param_2 := obj.DeQueue()
	param_3 := obj.Front()
	param_4 := obj.Rear()
	param_5 := obj.IsEmpty()
	param_6 := obj.IsFull()

	fmt.Println(param_1, param_2, param_3, param_4, param_5, param_6)
}

type MyCircularQueue struct {
	head  int
	tail  int
	queue []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	q := make([]int, k)
	circleQueue := MyCircularQueue{-1, -1, q}
	return circleQueue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.head == -1 && this.tail == -1 {
		this.head = 0
		this.tail = 0
		this.queue[this.tail] = value
		return true
	} else {
		if this.IsFull() {
			return false
		} else {
			this.tail += 1
			this.queue[this.tail%len(this.queue)] = value
			return true
		}
		return false
	}
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.head == -1 && this.tail == -1 {
		return false
	} else {
		if l := len(this.queue); this.tail%l == this.head%l {
			this.tail = -1
			this.head = -1
			return true
		} else {
			this.head += 1
			return true
		}
	}
	return false

}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.tail == -1 && this.head == -1 {
		return -1
	}
	return this.queue[this.head%len(this.queue)]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.tail == -1 && this.head == -1 {
		return -1
	}
	return this.queue[this.tail%len(this.queue)]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.tail == -1 && this.head == -1 {
		return true
	} else {
		return false
	}
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if l := len(this.queue); this.tail-this.head == l-1 {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
