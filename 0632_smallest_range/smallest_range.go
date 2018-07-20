package leetcode0632

import "container/heap"

func smallestRange(nums [][]int) []int {
	n := len(nums)
	h := &itemHeap{}

	arrPtrs := make([]int, len(nums))

	max := 0
	index := 0
	for index < n {
		arr := nums[index]
		arrPtr := arrPtrs[index]
		v := arr[arrPtr]
		if v > max {
			max = v
		}
		heap.Push(h, item{val: v, kth: index})
		arrPtrs[index] = arrPtr + 1
		index++
	}

	minDiff := max - (*h)[0].val
	res := []int{(*h)[0].val, max}

	for {
		minPop := heap.Pop(h).(item)
		arr := nums[minPop.kth]
		arrPtr := arrPtrs[minPop.kth]
		if arrPtr >= len(arr) {
			break
		}
		v := arr[arrPtr]
		if v > max {
			max = v
		}
		heap.Push(h, item{val: v, kth: minPop.kth})
		arrPtrs[minPop.kth] = arrPtr + 1

		curDiff := max - (*h)[0].val
		if curDiff < minDiff {
			minDiff = curDiff
			res = []int{(*h)[0].val, max}
		}
	}
	return res
}

type item struct {
	val int
	kth int
}

type itemHeap []item

func (q itemHeap) Less(i, j int) bool { return q[i].val < q[j].val }
func (q itemHeap) Len() int           { return len(q) }
func (q itemHeap) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *itemHeap) Push(x interface{}) {
	val := x.(item)
	*q = append(*q, val)
}

func (q *itemHeap) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
