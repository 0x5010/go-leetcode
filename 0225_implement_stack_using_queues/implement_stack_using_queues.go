package leetcode0225

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) Pop() int {
	num := q.nums[0]
	q.nums = q.nums[1:]
	return num
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

type MyStack struct {
	queue *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue: NewQueue()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue.Push(x)
	for i := 0; i < this.queue.Len()-1; i++ {
		this.queue.Push(this.queue.Pop())
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.queue.Pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	num := this.Pop()
	this.Push(num)
	return num
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
