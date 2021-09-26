package stack

import "time"

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value      string
		prev       *node
		time_stamp time.Time
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// Pop the top items of the stack and return it
func (this *Stack) Pop() (string, time.Time) {
	n := this.top
	this.top = n.prev
	this.length--
	return n.value, n.time_stamp
}

// Push items onto the top of the stack
func (this *Stack) Push(value string, time_stamp time.Time) {
	n := &node{value, this.top, time_stamp}
	this.top = n
	this.length++
}
