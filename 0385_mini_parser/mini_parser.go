package leetcode0385

import "strconv"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}

	if s[0] != '[' {
		return getValue(s)
	}

	stack := []*NestedInteger{}

	var res *NestedInteger
	i, j := 0, 0
	for j < len(s) {
		switch s[j] {
		case '[':
			if res != nil {
				stack = append(stack, res)
			}
			res = &NestedInteger{}
			i = j + 1
		case ']':
			if j > i {
				res.Add(*getValue(s[i:j]))
			}

			if len(stack) != 0 {
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmp.Add(*res)
				res = tmp
			}

			i = j + 1
		case ',':
			if s[j-1] != ']' {
				res.Add(*getValue(s[i:j]))
			}
			i = j + 1
		}
		j++
	}
	return res
}

func getValue(s string) *NestedInteger {
	val, _ := strconv.Atoi(s)
	ni := &NestedInteger{}
	ni.SetInteger(val)
	return ni
}
