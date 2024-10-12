import heapq

class Solution:
    def minGroups(self, intervals: List[List[int]]) -> int:
        # Sort intervals by start time
        intervals.sort()
        
        # Min-heap to track the end times of active groups
        pq = []
        
        # Traverse through all intervals
        for start, end in intervals:
            # If the earliest end time is less than the current start, reuse that group
            if pq and pq[0] < start:
                heapq.heappop(pq)
            
            # Add the current interval's end time to the heap
            heapq.heappush(pq, end)
        
        # The size of the heap is the number of groups
        return len(pq)
