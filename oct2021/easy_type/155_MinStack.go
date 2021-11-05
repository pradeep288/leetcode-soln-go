// https://www.youtube.com/watch?v=gd9xEAnxXzc&ab_channel=TECHDOSE
package easy_type

type MinStack struct {
	stack     []int
	min_stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		min_stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min_stack) == 0 {
		this.min_stack = append(this.min_stack, val)
	} else if this.min_stack[len(this.min_stack)-1] > val {
		this.min_stack = append(this.min_stack, val)
	} else {
		this.min_stack = append(this.min_stack, this.min_stack[len(this.min_stack)-1])
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min_stack = this.min_stack[:len(this.min_stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min_stack[len(this.min_stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
