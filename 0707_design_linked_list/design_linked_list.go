package leetcode0707

type MyLinkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	next *node
	val  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	n := &node{}
	return MyLinkedList{
		head: n,
		tail: n,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.size {
		return -1
	}
	p := this.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tmp := &node{
		next: this.head.next,
		val:  val,
	}
	if this.head == this.tail {
		this.tail = tmp
	}
	this.head.next = tmp
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tmp := &node{
		val: val,
	}
	this.tail.next = tmp
	this.tail = tmp
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	p := this.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	tmp := &node{
		next: p.next,
		val:  val,
	}
	if p == this.tail {
		this.tail = tmp
	}
	p.next = tmp
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	p := this.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	if p.next == this.tail {
		this.tail = p
	}
	p.next = p.next.next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
