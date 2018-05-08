package leetcode0385

type NestedInteger struct {
	Num int
	Ns  []*NestedInteger
}

func (n NestedInteger) IsInteger() bool { return n.Ns == nil }

func (n NestedInteger) GetInteger() int { return n.Num }

func (n *NestedInteger) SetInteger(value int) { n.Num = value }

func (n *NestedInteger) Add(elem NestedInteger) { n.Ns = append(n.Ns, &elem) }

func (n NestedInteger) GetList() []*NestedInteger { return n.Ns }
