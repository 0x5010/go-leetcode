package leetcode0232

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() int {
	num := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return num
}

func (s *Stack) Top() int {
	return s.nums[len(s.nums)-1]
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

type MyQueue struct {
	stack *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{stack: NewStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.Empty() {
		this.stack.Push(x)
		return
	}
	data := this.stack.Pop()
	this.Push(x)
	this.stack.Push(data)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.stack.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack.Top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
