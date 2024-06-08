# The Heap Property
# The top of the tree is always the min

# Note: This is a min-heap implementation


class Heap:
    def __init__(self) -> None:
        self.heap = []

    def parent(self, i):
        return (i - 1) // 2

    def l_child(self, i):
        return (2 * i) + 1

    def r_child(self, i):
        return (2 * i) + 2

    def add(self, num):
        # Add an element at the end of the array, and bubble up
        self.heap.append(num)
        self.up_heapify(len(self.heap) - 1)

    def extract(self):
        if len(self.heap) == 0:
            return None
        if len(self.heap) == 1:
            return self.heap.pop()

        min_element = self.heap[0]
        self.heap[0] = self.heap.pop()
        self.down_heapify(0)
        return min_element

    def up_heapify(self, i):
        while i != 0 and self.heap[self.parent(i)] > self.heap[i]:
            self.heap[self.parent(i)], self.heap[i] = (
                self.heap[i],
                self.heap[self.parent(i)],
            )
            self.up_heapify(self.parent(i))

    def down_heapify(self, i):
        # when extracting a min from the top, we need to replace the root
        # with the lowest element, and heapify_down, such that we can maintain
        # the heap property
        smallest = i
        left = self.l_child(i)
        right = self.r_child(i)
        heap_size = len(self.heap)

        # Check if left child exists and is smaller than top
        if left < heap_size and self.heap[left] < self.heap[smallest]:
            smallest = left

        if right < heap_size and self.heap[right] < self.heap[smallest]:
            smallest = right

        if smallest != i:
            self.heap[smallest], self.heap[i] = self.heap[i], self.heap[smallest]
            self.down_heapify(smallest)


heap = Heap()
heap.add(3)
heap.add(2)
heap.add(15)
heap.add(5)
heap.add(4)
heap.add(45)
print(heap.heap)
print(heap.extract())
print(heap.heap)
