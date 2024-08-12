import heapq

class KthLargest:
    def __init__(self, k: int, nums: list[int]):
        # Initialize the class with the desired 'k' and a list of numbers 'nums'
        self.k = k
        # Create an empty list to act as a min-heap
        self.minHeap = []
        
        # Process each number in the initial list 'nums'
        for num in nums:
            # Add each number to the heap using the add method
            self.add(num)

    def add(self, val: int) -> int:
        # If the heap has fewer than 'k' elements, push the new value onto the heap
        if len(self.minHeap) < self.k:
            heapq.heappush(self.minHeap, val)
        # If the heap already has 'k' elements, check if the new value is larger than the smallest element in the heap
        elif val > self.minHeap[0]:
            # If so, replace the smallest element with the new value
            heapq.heapreplace(self.minHeap, val)
        
        # Return the smallest element in the heap, which is the kth largest element in the stream
        return self.minHeap[0]
