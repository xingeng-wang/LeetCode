package Design

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
// push(x) -- Push element x onto stack.
// pop() -- Removes the element on top of the stack.
// top() -- Get the top element.
// getMin() -- Retrieve the minimum element in the stack.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/562/
type MinStack struct {
	val    []int
	length int
	min    *int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{val: []int{}, min: nil, length: 0}
}

func (this *MinStack) Push(x int) {
	this.val = append(this.val, x)
	this.length += 1
	if this.min == nil || x < *this.min {
		this.min = &this.val[this.length-1]
	}
}

func (this *MinStack) Pop() {
	if this.length > 0 {
		this.val = this.val[:this.length-1]
		this.length = this.length - 1
		if this.length == 0 {
			this.min = nil
			return
		}
		this.min = &this.val[0]
		for i, n := range this.val {
			if n < *this.min {
				this.min = &this.val[i]
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.val[this.length-1]
}

func (this *MinStack) GetMin() int {
	return *this.min
}
