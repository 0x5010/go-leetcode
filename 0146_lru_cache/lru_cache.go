package leetcode0146

type dLinkedNode struct {
	pre, post  *dLinkedNode
	key, value int
}

type LRUCache struct {
	len, cap   int
	head, tail *dLinkedNode
	nodes      map[int]*dLinkedNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		nodes: make(map[int]*dLinkedNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodes[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodes[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}

	if this.len == this.cap {
		delete(this.nodes, this.tail.key)
		this.pop()
	} else {
		this.len++
	}
	node := &dLinkedNode{
		key:   key,
		value: value,
	}
	this.nodes[key] = node
	this.add(node)
}

func (this *LRUCache) pop() {
	if this.tail.pre != nil {
		this.tail.pre.post = nil
	} else {
		this.head = nil
	}
	this.tail = this.tail.pre
}

func (this *LRUCache) add(node *dLinkedNode) {
	if this.tail == nil {
		this.tail = node
	} else {
		node.post = this.head
		this.head.pre = node
	}
	this.head = node
}

func (this *LRUCache) moveToHead(node *dLinkedNode) {
	if node == this.head {
		return
	} else if node == this.tail {
		this.pop()
	} else {
		node.pre.post = node.post
		node.post.pre = node.pre
	}
	this.add(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
