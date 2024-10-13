import heapq

class Solution:
    def smallestRange(self, nums: List[List[int]]) -> List[int]:
        minHeap = []
        maxValue = float('-inf')
        
        # Initialize heap with the first element from each list
        for i in range(len(nums)):
            heapq.heappush(minHeap, (nums[i][0], i, 0))
            maxValue = max(maxValue, nums[i][0])
        
        rangeStart, rangeEnd = 0, float('inf')
        
        while minHeap:
            minValue, row, col = heapq.heappop(minHeap)
            
            # Update the smallest range
            if maxValue - minValue < rangeEnd - rangeStart:
                rangeStart, rangeEnd = minValue, maxValue
            
            # Move to the next element in the current list
            if col + 1 < len(nums[row]):
                nextValue = nums[row][col + 1]
                heapq.heappush(minHeap, (nextValue, row, col + 1))
                maxValue = max(maxValue, nextValue)
            else:
                break  # One list is exhausted
        
        return [rangeStart, rangeEnd]