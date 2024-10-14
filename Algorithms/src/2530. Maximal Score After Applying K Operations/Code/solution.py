import heapq
import math

class Solution:
    def maxKelements(self, nums: List[int], k: int) -> int:
        # Max-heap (we negate values to simulate max-heap behavior)
        maxHeap = [-num for num in nums]
        heapq.heapify(maxHeap)
        
        score = 0
        
        # Perform k operations
        for _ in range(k):
            # Get the largest element (by negating to retrieve the max)
            maxVal = -heapq.heappop(maxHeap)
            
            # Add to the score
            score += maxVal
            
            # Replace the element with ceil(maxVal / 3)
            heapq.heappush(maxHeap, -math.ceil(maxVal / 3))
        
        return score