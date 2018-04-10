package leetcode0307

type Node struct {
	l, r *Node
	i, j int
	sum  int
}

func buildTree(nums []int, i, j int) *Node {
	if i > j {
		return nil
	}
	node := &Node{
		i: i,
		j: j,
	}
	if i == j {
		node.sum = nums[i]
	} else {
		m := (i + j) / 2
		node.l = buildTree(nums, i, m)
		node.r = buildTree(nums, m+1, j)
		node.sum = node.l.sum + node.r.sum
	}
	return node
}

func (node *Node) update(i, val int) {
	if node.i == node.j {
		node.sum = val
		return
	}
	m := (node.i + node.j) / 2
	if i <= m {
		node.l.update(i, val)
	} else {
		node.r.update(i, val)
	}
	node.sum = node.l.sum + node.r.sum
}

func (node *Node) sumRange(i, j int) int {
	if node.i == i && node.j == j {
		return node.sum
	}
	m := (node.i + node.j) / 2
	if j <= m {
		return node.l.sumRange(i, j)
	} else if i >= m+1 {
		return node.r.sumRange(i, j)
	} else {
		return node.l.sumRange(i, m) + node.r.sumRange(m+1, j)
	}
}

type NumArray struct {
	root *Node
}

func Constructor(nums []int) NumArray {
	return NumArray{
		root: buildTree(nums, 0, len(nums)-1),
	}
}

func (this *NumArray) Update(i int, val int) {
	this.root.update(i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.root.sumRange(i, j)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
