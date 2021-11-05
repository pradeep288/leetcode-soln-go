package easy_type

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
}

func (this *MyStack) Pop() int {
	val := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return val
}

func (this *MyStack) Top() int {
	return this.q[len(this.q)-1]
}

func (this *MyStack) Empty() bool {
	return this.q[len(this.q)] == 0
}
