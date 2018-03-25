package leetcode0215

type heap struct {
	heap []int
}

func (h *heap) getParent(child int) int {
	if child%2 == 0 {
		return (child / 2) - 1
	}
	return child / 2

}

func (h *heap) getLeftChild(node int) int {
	return node*2 + 1
}

func (h *heap) getRightChild(node int) int {
	return (node + 1) * 2
}

func (h *heap) remove() {
	if len(h.heap) <= 2 {
		h.heap = h.heap[1:]
		return
	}
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	parent := 0

	for true {
		l := h.getLeftChild(parent)
		r := h.getRightChild(parent)

		largest := parent

		if l < len(h.heap) && h.heap[largest] > h.heap[l] {
			largest = l
		}
		if r < len(h.heap) && h.heap[largest] > h.heap[r] {
			largest = r
		}

		if parent != largest {
			h.heap[parent], h.heap[largest] = h.heap[largest], h.heap[parent]
			parent = largest
		} else {
			break
		}
	}
}

func (h *heap) insert(val int) {
	h.heap = append(h.heap, val)

	child := len(h.heap) - 1
	parent := h.getParent(child)

	for child != 0 && h.heap[child] < h.heap[parent] {
		h.heap[child], h.heap[parent] = h.heap[parent], h.heap[child]
		child = parent
		parent = h.getParent(child)
	}
}

func (h *heap) heapify(nums []int) {
	for _, val := range nums {
		h.insert(val)
	}
}

func (h *heap) min() int {
	if len(h.heap) > 0 {
		return h.heap[0]
	}
	return -1
}

func findKthLargest(nums []int, k int) int {
	h := &heap{heap: []int{}}
	h.heapify(nums[:k])

	for i := k; i < len(nums); i++ {
		if nums[i] > h.min() {
			h.remove()
			h.insert(nums[i])
		}
	}
	return h.min()
}
