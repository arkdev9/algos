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
        current = i
        while current != 0 and self.heap[self.parent(current)] > self.heap[current]:
            self.heap[self.parent(current)], self.heap[current] = (
                self.heap[current],
                self.heap[self.parent(current)],
            )
            current = self.parent(current)

    def down_heapify(self, i):
        smallest = i
        left = self.l_child(i)
        right = self.r_child(i)
        heap_size = len(self.heap)

        if left < heap_size and self.heap[left] < self.heap[smallest]:
            smallest = left
        if right < heap_size and self.heap[right] < self.heap[smallest]:
            smallest = right

        if smallest != i:
            self.heap[smallest], self.heap[i] = self.heap[i], self.heap[smallest]
            self.down_heapify(smallest)


# Test the heap operations
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
print(heap.extract())
print(heap.heap)
