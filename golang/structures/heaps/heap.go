package heap

// Heap operations
// - create a heap (make and return an empty integer slice)

var length = 10

func CreateHeap() []int {
	return make([]int, 5)
}

func HeapifyUp(idx int, heap []int) []int {
	// Heapify up starts bubbling up given index with parent if possible
	// Given an idx, compare with parent, if parent is greater than currValue
	// then swap parentValue wtih currValue (essentially swapping node).
	// recurse with the parent idx if swap has taken place
	if idx == 0 {
		return nil
	}

	parentValue := heap[parent(idx, heap)]
	value := heap[idx]

	if value < parentValue {
		heap[parent(idx, heap)] = value
		heap[idx] = parentValue
		HeapifyUp()
	}

	return make([]int, 5)
}

func parent(idx int, heap []int) int {
	// Parent will always be
	return (len(heap) - 1) / 2
}

func leftChild(idx int, heap []int) int {
	return (2*idx + 1)
}

func rightChild(idx int, heap []int) int {
	return (2*idx + 2)
}
