package leetcode0919

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type CBTInserter struct {
	tree []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	tree := []*TreeNode{root}
	for i := 0; i < len(tree); i++ {
		if tree[i].Left != nil {
			tree = append(tree, tree[i].Left)
		}
		if tree[i].Right != nil {
			tree = append(tree, tree[i].Right)
		}
	}
	return CBTInserter{
		tree: tree,
	}
}

func (this *CBTInserter) Insert(v int) int {
	n := len(this.tree)
	node := &TreeNode{Val: v}
	this.tree = append(this.tree, node)
	index := (n - 1) / 2
	if n%2 == 0 {
		this.tree[index].Right = node
	} else {
		this.tree[index].Left = node
	}
	return this.tree[index].Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	if len(this.tree) != 0 {
		return this.tree[0]
	}
	return nil
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
